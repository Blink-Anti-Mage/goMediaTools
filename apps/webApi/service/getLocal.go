package service

import (
	"fmt"
	"goMediatools/datacache"
	"goMediatools/internal/config"
	"goMediatools/model"
	"os"
	"path/filepath"
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
