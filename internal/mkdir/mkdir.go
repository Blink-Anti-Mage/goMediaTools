package mkdir

import (
	"fmt"
	"os"
)

func Mkdir(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Redir(dirName string) {
	err := os.Remove(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
