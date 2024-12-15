package config

import "github.com/BurntSushi/toml"

type Config struct {
	Database struct {
		Host     string `toml:"DB_HOST"`
		Port     string `toml:"DB_PORT"`
		User     string `toml:"DB_USER"`
		Password string `toml:"DB_PASSWORD"`
		Name     string `toml:"DB_NAME"`
		SSLMode  string `toml:"DB_SSLMODE"`
	}
	JWT struct {
		AccessSecret  string `toml:"JWT_ACCESS_SECRET"`
		RefreshSecret string `toml:"JWT_REFRESH_SECRET"`
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
