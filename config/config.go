package config

import (
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Server    ServerConfig
	Redis     RedisConfig
	URLPrefix string `env:"URL_PREFIX" envDefault:"http://localhost:8080/"`
}

type ServerConfig struct {
	Addr         string        `env:"ADDR" envDefault:":8080"`
	ReadTimeout  time.Duration `env:"TIMEOUT_READ" envDefault:"5s"`
	WriteTimeout time.Duration `env:"TIMEOUT_WRITE" envDefault:"10s"`
	IdleTimeout  time.Duration `env:"TIMEOUT_IDLE" envDefault:"15s"`
}

type RedisConfig struct {
	Addr       string        `env:"REDIS_ADDR"`
	Password   string        `env:"REDIS_PASSWORD"`
	Expiration time.Duration `env:"REDIS_EXPIRATION" envDefault:"0"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.Server); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.Redis); err != nil {
		return nil, err
	}

	return &cfg, nil
}
