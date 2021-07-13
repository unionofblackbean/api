package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	App *AppConfig `json:"app"`
}

func Parse(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json from bytes -> %v", err)
	}

	return &config, nil
}

func ParseString(content string) (*Config, error) {
	return Parse([]byte(content))
}

func ParseFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file -> %v", err)
	}

	parser := NewParser(file)
	config, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from reader -> %v", err)
	}

	return config, nil
}
