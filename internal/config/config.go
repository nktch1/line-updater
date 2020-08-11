package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Request struct {
		UpdateIntervalFootball int `envconfig:"UPD_INTERVAL_FOOTBALL"`
		UpdateIntervalSoccer   int `envconfig:"UPD_INTERVAL_SOCCER"`
		UpdateIntervalBaseball int `envconfig:"UPD_INTERVAL_BASEBALL"`
	}
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Database struct {
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
	}
	Log struct {
		Level string `envconfig:"LOG_LEVEL"`
	}
	LineProvider struct {
		URL string `envconfig:"LINE_PROVIDER_API_URL"`
	}
	RPCServer struct {
		Host string `envconfig:"RPC_SERVER_HOST"`
		Port string `envconfig:"RPC_SERVER_PORT"`
	}
}

// New загружает значения переменных среды
func New() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		//cfg.Request.UpdateIntervalBaseball = 12
		//cfg.Request.UpdateIntervalFootball = 12
		//cfg.Request.UpdateIntervalSoccer = 12
		//cfg.Database.Username = "user"
		//cfg.Database.Password = "123"
		//cfg.Server.Host = "localhost"
		//cfg.Server.Port = "8080"

		//return &cfg, fmt.Errorf("\tconfig file parsing failed, | [%s]", err.Error())
		return &cfg
	}

	return &cfg
}
