package config

import "github.com/BurntSushi/toml"

type Config struct {
	Hosts struct {
		TextgenService string `toml:"TEXTGEN_SERVICE_HOST"`
		ApiService     string `toml:"API_SERVICE_HOST"`
	}
	Logger struct {
		LogLevel string `toml:"LOG_LEVEL"`
	}
}

var Data Config

func LoadConfig(path string) error {
	_, err := toml.DecodeFile(path, &Data)
	if err != nil {
		return err
	}
	return nil
}
