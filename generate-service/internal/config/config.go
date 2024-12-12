package config

import "github.com/BurntSushi/toml"

type Config struct {
	TextgenServiceHost string `toml:"TEXTGEN_SERVICE_HOST"`
	ApiServiceHost     string `toml:"API_SERVICE_HOST"`
}

var Data Config

func LoadConfig(path string) error {
	_, err := toml.DecodeFile(path, &Data)
	if err != nil {
		return err
	}
	return nil
}
