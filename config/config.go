package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Rest RestConfig `json:"rest"`
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
