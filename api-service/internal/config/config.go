package config

type Config struct {
	Database struct {
		Host     string `toml:"db_host"`
		Port     string `toml:"db_port"`
		User     string `toml:"db_user"`
		Password string `toml:"db_password"`
		Name     string `toml:"db_name"`
		SSLMode  string `toml:"db_sslmode"`
	}
	JWT struct {
		AccessSecret  string `toml:"jwt_access_secret"`
		RefreshSecret string `toml:"jwt_refresh_secret"`
	}
	Logger struct {
		LogLevel string `toml:"log_level"`
	}
	ApiServer struct {
		BindAddr string `toml:"bind_addr"`
	}
	GrpcServer struct {
		BindAddr string `toml:"bind_addr"`
	}
}

func NewConfig() *Config {
	return &Config{}
}
