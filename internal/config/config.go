package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Con Config

func InitConfig(filepath string) error {
	fmt.Println("add config" + filepath)
	if filepath == "" {
		filepath = "./config.json"
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Con)
	if err != nil {
		return err
	}
	fmt.Println(Con)
	return nil
}

// config
type Config struct {
	SuffixList []string `json:"SuffixList"`
	MovieDir   string   `json:"MovieDir"`
}
