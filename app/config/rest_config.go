package config

type RestConfig struct {
	BindAddr string `json:"bind_addr"`
	BindPort uint16 `json:"bind_port"`

	RateLimit int `json:"rate_limit"`
}
