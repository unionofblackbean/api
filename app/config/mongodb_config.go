package config

type MongoDBConfig struct {
	Addr string `json:"addr"`
	Port uint16 `json:"port"`

	Username string `json:"username"`
	Password string `json:"password"`

	Timeout int64 `json:"timeout"`
}
