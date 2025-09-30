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
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func New(configPath string) (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
