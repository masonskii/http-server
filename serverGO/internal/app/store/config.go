package store

//Config ...
type Config struct {
	dbURL string `toml:"database_url"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
