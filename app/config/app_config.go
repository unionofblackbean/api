package config

type AppConfig struct {
	StartPolicy StartPolicy `json:"start_policy"`

	DB *DBConfig `json:"db"`

	Services *ServicesConfig `json:"services"`
}
