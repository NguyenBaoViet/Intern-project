package conf

import "github.com/caarlos0/env"

type AppConfig struct {
	Port      string `env:"PORT" envDefault:""`
	LogFormat string `env:"LOG_FORMAT" envDefault:""`

	DBHost string `env:"DB_HOST" envDefault:""`
	DBPort string `env:"DB_PORT" envDefault:""`
	DBUser string `env:"DB_USER" envDefault:""`
	DBPass string `env:"DB_PASS" envDefault:""`
	DBName string `env:"DB_NAME" envDefault:""`

	EnableDB string `env:"ENABLE_DB" envDefault:""`

	SecretKey string `env:"SECRET_KEY"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
