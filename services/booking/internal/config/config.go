package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host string `yaml:"host" env:"HOST" env-required:"true"`
	Port int    `yaml:"port" env:"PORT" env-required:"true"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"     env:"HOST"     env-required:"true"`
	Port     int    `yaml:"port"     env:"PORT"     env-required:"true"`
	User     string `yaml:"user"     env:"USER"     env-required:"true"`
	Password string `yaml:"password" env:"PASSWORD" env-required:"true"`
	DB       string `yaml:"db"       env:"DB"       env-required:"true"`
	SSLMode  string `yaml:"sslmode"  env:"SSLMODE"  env-required:"true"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"   env-prefix:"SERVER_"`
	Postgres PostgresConfig `yaml:"postgres" env-prefix:"POSTGRES_"`
}

func New(configPath string) (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
