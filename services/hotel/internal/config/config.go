package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	SSLMode  string `yaml:"sslmode"`
	Port     int    `yaml:"port"`
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Env      string         `yaml:"env"`
	LogLevel string         `yaml:"log_level"`
	Server   ServerConfig   `yaml:"server"`
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
