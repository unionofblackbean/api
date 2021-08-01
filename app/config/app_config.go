package config

type AppConfig struct {
	StartPolicy StartPolicy `json:"start_policy"`

	Postgres *PostgresConfig `json:"postgres"`
	MongoDB  *MongoDBConfig  `json:"mongodb"`

	Services *ServicesConfig `json:"services"`
}
