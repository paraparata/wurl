package config

type Config struct {
	Store *[]byte
}

func New(store *[]byte) *Config {
	return &Config{store}
}
