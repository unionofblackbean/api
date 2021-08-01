package config

type PostgresConfig struct {
	Addr string `json:"addr"`
	Port uint16 `json:"port"`

	Username string `json:"username"`
	Password string `json:"password"`

	DBName string `json:"db_name"`

	Timeout int64 `json:"timeout"`
}
