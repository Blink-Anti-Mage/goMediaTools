package service

import (
	"encoding/json"
	"fmt"
	"goMediatools/internal/config"
	"testing"
)

func TestGetlocal(t *testing.T) {
	err := config.InitConfig("D:\\code\\golang_code\\goMediaTools\\config.json")
	if err != nil {
		fmt.Println("add config err:" + err.Error())
		return
	}

	Getlocal()
	// jsonData, err := json.Marshal(datacache.LocalCache)
	// if err != nil {
	// 	fmt.Printf("JSON marshaling failed: %s", err)
	// }
	// fmt.Println(string(jsonData))
}
func TestGetlocalF(t *testing.T) {
	rootPath := `\\nas\media\dy`
	tree, err := BuildTree(rootPath)
	if err != nil {
	}
	treeJSON, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
	}
	fmt.Println(string(treeJSON))
}
func TestMove(t *testing.T) {
	rootPath := `\\nas\media\dy`
	err := MoveDir(rootPath)
	if err != nil {
	}
}
func TestMove2(t *testing.T) {
	err := config.InitConfig("D:\\code\\golang_code\\goMediaTools\\config.json")
	if err != nil {
		fmt.Println("add config err:" + err.Error())
		return
	}

	rootPath := `\\nas\media\dy`
	err = MoveDir2(rootPath)
	if err != nil {

	}
}
