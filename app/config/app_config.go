package config

type AppConfig struct {
	StartPolicy StartPolicy `json:"start_policy"`

	Postgres *PostgresConfig `json:"postgres"`
	Mongo    *MongoConfig    `json:"mongo"`

	Services *ServicesConfig `json:"services"`
}
