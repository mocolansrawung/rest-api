package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database struct {
		DBName string `json:"DBName"`
	} `json:"Database"`
	Server struct {
		Port string `json:"Port"`
	} `json:"Server"`
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
