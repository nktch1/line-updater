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
		cfg.Request.UpdateIntervalBaseball = 1
		cfg.Request.UpdateIntervalFootball = 2
		cfg.Request.UpdateIntervalSoccer = 3
		cfg.Database.Username = ""
		cfg.Database.Password = ""
		cfg.Server.Host = "localhost"
		cfg.Server.Port = "8080"
		cfg.Log.Level = "debug"
		cfg.RPCServer.Host = ""
		cfg.RPCServer.Port = "8888"
		cfg.Database.Host = "redis"
		cfg.Database.Port = "6379"
		cfg.LineProvider.URL = "http://lineprovider:8000/api/v1/lines"
	}

	return &cfg
}
