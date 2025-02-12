package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database   DatabaseConfig
	JWT        JWTConfig
	Logger     LoggerConfig
	ApiServer  ApiServerConfig
	GrpcServer GrpcServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
}

type LoggerConfig struct {
	LogLevel string
}

type ApiServerConfig struct {
	BindAddr string
}

type GrpcServerConfig struct {
	BindAddr string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	fmt.Println(config)
	return &config, nil
}
