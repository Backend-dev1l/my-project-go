package logger

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	DevMode  = "dev"
	ProdMode = "prod"
)

type Option func(*settings)

type settings struct {
	writer io.Writer
}

type Config struct {
	Mode  string `yaml:"mode" env:"MODE" env-default:"dev" validate:"oneof=dev prod"`
	Level string `yaml:"level" env:"LEVEL" env-default:"info" validate:"oneof=debug info warn warning error"`
}

func WithWriter(w io.Writer) Option {
	return func(s *settings) { s.writer = w }
}

func defaultSettings() settings {
	return settings{writer: os.Stdout}
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `yaml:"logger" env-prefix:"DB_"`
	}
	if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		if err = cleanenv.ReadEnv(&cfg); err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}
func New(cfg *Config, opts ...Option) (*slog.Logger, error) {
	if cfg == nil {
		return nil, errors.New("logger: nil config")
	}

	s := defaultSettings()
	for _, o := range opts {
		o(&s)
	}

	lvl := parseLevel(cfg.Level)

	var h slog.Handler
	switch cfg.Mode {
	case DevMode:
		h = slog.NewTextHandler(s.writer, &slog.HandlerOptions{Level: lvl})
	case ProdMode:
		h = slog.NewJSONHandler(s.writer, &slog.HandlerOptions{Level: lvl})
	default:
		return nil, fmt.Errorf("unknown logger mode: %s", cfg.Mode)
	}

	return slog.New(h), nil
}

func Init(l *slog.Logger) {

	slog.SetDefault(l)

}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
