package service

import (
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
