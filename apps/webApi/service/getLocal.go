package service

import (
	"fmt"
	"goMediatools/datacache"
	"goMediatools/internal/config"
	"goMediatools/model"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Getlocal() {
	path := config.Con.MovieDir
	if config.Con.MovieDir != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("path not exist:", path)
		} else if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println("path exist:", path)
		}
	}

	err := filepath.Walk(config.Con.MovieDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			_, filen := filepath.Split(path)
			fileame := filen[:len(filen)-len(filepath.Ext(filen))]
			for _, suffix := range config.Con.SuffixList {
				if ext == suffix {
					chNfo, chPoster, chPath, poster := CheckNfo(path)
					datacache.LocalCache[fileame] = model.Movie{Name: fileame, Path: path, IfNfo: chNfo,
						IfPoster: chPoster, IfPath: chPath, PosterDir: poster}
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path err")
	}

}

func CheckNfo(path string) (nfo bool, poster bool, chpath bool, posterdir string) {
	nfo = false
	poster = false
	chpath = false
	posterdir = ""

	dir := filepath.Dir(filepath.Dir(path))
	if dir != config.Con.MovieDir {
		chpath = false
	} else {
		chpath = true
	}

	dir, filen := filepath.Split(path)
	fileame := filen[:len(filen)-len(filepath.Ext(filen))]
	nfoname := fileame + ".nfo"
	jpgposter := fileame + "-poster.jpg"
	pngposter := fileame + "-poster.png"

	_, err := os.Stat(filepath.Join(dir, nfoname))
	if os.IsNotExist(err) {
		nfo = false
	} else {
		nfo = true
	}

	_, err = os.Stat(filepath.Join(dir, jpgposter))
	if os.IsNotExist(err) {
		_, err = os.Stat(filepath.Join(dir, pngposter))
		if os.IsNotExist(err) {
			poster = false
		} else {
			posterdir = filepath.Base(dir) + "/" + pngposter

			poster = true
		}
	} else {
		posterdir = filepath.Base(dir) + "/" + jpgposter
		poster = true
	}
	return nfo, poster, chpath, posterdir
}

// TreeNode 代表一个文件树节点
type TreeNode struct {
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Ifmovie  bool       `json:"ifmovie"`
	Ifnfo    bool       `json:"ifnfo"`
	Children []TreeNode `json:"children,omitempty"`
}

func BuildTree(rootPath string) (TreeNode, error) {
	rootNode := TreeNode{Name: filepath.Base(rootPath),
		Path: rootPath, Ifmovie: false, Ifnfo: false,
		Children: []TreeNode{}}

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return rootNode, err
	}

	for _, file := range files {
		fullPath := filepath.Join(rootPath, file.Name())
		if file.IsDir() {
			childNode, err := BuildTree(filepath.Join(rootPath, file.Name()))
			if err != nil {
				return rootNode, err
			}
			rootNode.Children = append(rootNode.Children, childNode)
		} else {
			ifmovie := false
			ifnfo := false
			if strings.HasSuffix(file.Name(), ".nfo") {
				ifnfo = true
			} else {
				for _, suffix := range config.Con.SuffixList {
					if strings.HasSuffix(file.Name(), suffix) {
						ifmovie = true
					}
				}
			}
			rootNode.Children = append(rootNode.Children, TreeNode{Name: file.Name(),
				Path: fullPath, Ifmovie: ifmovie, Ifnfo: ifnfo})

		}
	}

	return rootNode, nil
}

func MoveDir(rootPath string) error {
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fullPath := filepath.Join(rootPath, file.Name())

			// 去掉文件后缀
			fileName := file.Name()
			ext := filepath.Ext(fileName)
			newDirName := strings.TrimSuffix(fileName, ext)

			// 创建新的文件夹
			newDirPath := filepath.Join(rootPath, newDirName)
			if err := os.Mkdir(newDirPath, os.ModePerm); err != nil {
				return err
			}

			// 移动文件到新的文件夹
			newFilePath := filepath.Join(newDirPath, fileName)
			if err := os.Rename(fullPath, newFilePath); err != nil {
				return err
			}

			fmt.Printf("Moved %s to %s\n", fullPath, newFilePath)
		}
	}
	return nil
}

func MoveDir2(rootPath string) error {
	var wg sync.WaitGroup
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			Path := filepath.Join(rootPath, file.Name())
			wg.Add(1)
			go func() {
				defer wg.Done()
				Renameworker(Path)
			}()
		}
	}
	wg.Wait()
	return nil
}

func Renameworker(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	ifmovie := ""
	ifnfo := ""
	ifposterjpg := ""
	ifposterpng := ""
	newDirName := ""

	for _, file := range files {
		if !file.IsDir() {
			for _, suffix := range config.Con.SuffixList {
				if strings.HasSuffix(file.Name(), suffix) {
					ifmovie = filepath.Join(path, file.Name())
					continue
				}

				if strings.HasSuffix(file.Name(), ".nfo") {
					ifnfo = filepath.Join(path, file.Name())
					continue
				}
				if strings.HasSuffix(file.Name(), "-poster.jpg") {
					ifposterjpg = filepath.Join(path, file.Name())
					continue
				}
				if strings.HasSuffix(file.Name(), "-poster.png") {
					ifposterpng = filepath.Join(path, file.Name())
					continue
				}
			}
		}
	}

	if ifnfo != "" {
		data, err := os.ReadFile(ifnfo)
		if err != nil {
			return
		}
		jsondata, err := GetNFO(data)
		if err != nil {
			fmt.Println("nfo ioread error")
			return
		}
		if jsondata.Title != "" {
			newDirName = jsondata.Title + "(" + jsondata.Year + ")"
		}
	}
	if newDirName != "" {
		newname := filepath.Join(path, newDirName+".nfo")
		if err := os.Rename(ifnfo, newname); err != nil {
			return
		}
		if ifposterjpg != "" {
			newname := filepath.Join(path, newDirName+"-poster.jpg")
			if err := os.Rename(ifposterjpg, newname); err != nil {
				return
			}
		}
		if ifposterpng != "" {
			newname := filepath.Join(path, newDirName+"-poster.png")
			if err := os.Rename(ifposterpng, newname); err != nil {
				return
			}
		}
		if ifmovie != "" {
			newname := filepath.Join(path, newDirName+filepath.Ext(ifmovie))
			if err := os.Rename(ifmovie, newname); err != nil {
				return
			}
		}

		if err := os.Rename(path, filepath.Join(filepath.Dir(path), newDirName)); err != nil {
			return
		}

	}

}
