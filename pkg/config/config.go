package config

type Config struct {
	Store    []byte
	PathFile string
}

func New(pathFile string) *Config {
	return &Config{PathFile: pathFile}
}
