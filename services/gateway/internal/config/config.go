package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCServiceConfig struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"PORT" env-required:"true"`
}

type HTTPConfig struct {
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"PORT" env-default:"8080"`
}

type JWTConfig struct {
	AccessSecret string `yaml:"access_secret" env:"ACCESS_SECRET" env-required:"true"`
}

type Config struct {
	Env     string         `yaml:"env" env:"ENV" env-default:"local"`
	LogLevel string        `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`
	HTTP    HTTPConfig     `yaml:"http" env-prefix:"HTTP_"`
	JWT     JWTConfig      `yaml:"jwt" env-prefix:"JWT_"`
	Auth    GRPCServiceConfig `yaml:"auth" env-prefix:"AUTH_"`
	Hotel   GRPCServiceConfig `yaml:"hotel" env-prefix:"HOTEL_"`
	Booking GRPCServiceConfig `yaml:"booking" env-prefix:"BOOKING_"`
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
