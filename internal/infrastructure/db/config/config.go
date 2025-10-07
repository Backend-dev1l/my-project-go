package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"password"`
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DB_NAME" env-default:"appdb"`
	SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `yaml:"database" env-prefix:"DB_"`
	}
	if err := cleanenv.ReadConfig("config/config.yaml", &cfg); err != nil {
		if err = cleanenv.ReadEnv(&cfg); err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode,
	)
}
