package config

type AppConfig struct {
	StartPolicy StartPolicy `json:"start_policy"`

	Postgres *PostgresConfig `json:"postgres"`

	Services *ServicesConfig `json:"services"`
}
