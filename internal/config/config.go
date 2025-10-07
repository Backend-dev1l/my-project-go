package config

import (
	"net"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int    `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	Host string `yaml:"host" env:"SERVER_HOST" env-default:"127.0.0.1"`
}

func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `yaml:"http" env-prefix:"HTTP_"`
	}
	if err := cleanenv.ReadConfig("config/config.yaml", &cfg); err != nil {
		if err = cleanenv.ReadEnv(&cfg); err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}
