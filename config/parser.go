package config

import (
	"encoding/json"
	"fmt"
	"io"
)

type Parser struct {
	dec *json.Decoder
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		dec: json.NewDecoder(reader),
	}
}

func (p *Parser) Parse() (*Config, error) {
	var config Config
	err := p.dec.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json from reader -> %v", err)
	}

	return &config, nil
}
