package config

import (
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Server  ServerConfig
	Service ServiceConfig
	Redis   RedisConfig
}

type ServerConfig struct {
	Host         string        `env:"HTTP_HOST,required"`
	Port         string        `env:"HTTP_PORT,required"`
	ReadTimeout  time.Duration `env:"TIMEOUT_READ" envDefault:"5s"`
	WriteTimeout time.Duration `env:"TIMEOUT_WRITE" envDefault:"10s"`
	IdleTimeout  time.Duration `env:"TIMEOUT_IDLE" envDefault:"15s"`
}

type ServiceConfig struct {
	URLPrefix string `env:"URL_PREFIX,required"`
}

type RedisConfig struct {
	Addr       string        `env:"REDIS_ADDR,required"`
	Password   string        `env:"REDIS_PASSWORD,required"`
	Expiration time.Duration `env:"REDIS_EXPIRATION,required"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.Server); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.Service); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.Redis); err != nil {
		return nil, err
	}

	return &cfg, nil
}
