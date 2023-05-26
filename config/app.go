package config

import (
	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	LogLevel     int    `env:"APP_LOG_LEVEL"`
	AppPort      int    `env:"APP_PORT"`
	AppProxyPort int    `env:"APP_PROXY_PORT"`
	CryptoKey    string `env:"CRYPTO_KEY"`
}

func NewAppConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
