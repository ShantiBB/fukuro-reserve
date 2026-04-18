package config

import "github.com/ilyakaznacheev/cleanenv"

type GRPCServiceConfig struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"PORT" env-required:"true"`
}

type HTTPConfig struct {
	Host               string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Port               int    `yaml:"port" env:"PORT" env-default:"8080"`
	HealthPath         string `yaml:"health_path" env:"HEALTH_PATH" env-default:"/health"`
	APIPrefix          string `yaml:"api_prefix" env:"API_PREFIX" env-default:"/api/v1"`
	ReadTimeoutSec     int    `yaml:"read_timeout_sec" env:"READ_TIMEOUT_SEC" env-default:"15"`
	WriteTimeoutSec    int    `yaml:"write_timeout_sec" env:"WRITE_TIMEOUT_SEC" env-default:"15"`
	IdleTimeoutSec     int    `yaml:"idle_timeout_sec" env:"IDLE_TIMEOUT_SEC" env-default:"60"`
	RequestTimeoutSec  int    `yaml:"request_timeout_sec" env:"REQUEST_TIMEOUT_SEC" env-default:"60"`
	ShutdownTimeoutSec int    `yaml:"shutdown_timeout_sec" env:"SHUTDOWN_TIMEOUT_SEC" env-default:"30"`
}

type PaginationConfig struct {
	DefaultPage     uint64 `yaml:"default_page" env:"DEFAULT_PAGE" env-default:"1"`
	DefaultPageSize uint64 `yaml:"default_page_size" env:"DEFAULT_PAGE_SIZE" env-default:"100"`
}

type CORSConfig struct {
	AllowedOrigins   []string `yaml:"allowed_origins" env:"ALLOWED_ORIGINS" env-separator:"," env-default:"*"`
	AllowedMethods   []string `yaml:"allowed_methods" env:"ALLOWED_METHODS" env-separator:"," env-default:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowedHeaders   []string `yaml:"allowed_headers" env:"ALLOWED_HEADERS" env-separator:"," env-default:"Accept,Authorization,Content-Type,X-CSRF-Token"`
	ExposedHeaders   []string `yaml:"exposed_headers" env:"EXPOSED_HEADERS" env-separator:"," env-default:"Link"`
	AllowCredentials bool     `yaml:"allow_credentials" env:"ALLOW_CREDENTIALS" env-default:"true"`
	MaxAge           int      `yaml:"max_age" env:"MAX_AGE" env-default:"300"`
}

type JWTConfig struct {
	AccessSecret string `yaml:"access_secret" env:"ACCESS_SECRET" env-required:"true"`
}

type Config struct {
	Env        string            `yaml:"env" env:"ENV" env-default:"local"`
	LogLevel   string            `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`
	HTTP       HTTPConfig        `yaml:"http" env-prefix:"HTTP_"`
	Pagination PaginationConfig  `yaml:"pagination" env-prefix:"PAGINATION_"`
	CORS       CORSConfig        `yaml:"cors" env-prefix:"CORS_"`
	JWT        JWTConfig         `yaml:"jwt" env-prefix:"JWT_"`
	Auth       GRPCServiceConfig `yaml:"auth" env-prefix:"AUTH_"`
	Hotel      GRPCServiceConfig `yaml:"hotel" env-prefix:"HOTEL_"`
	Booking    GRPCServiceConfig `yaml:"booking" env-prefix:"BOOKING_"`
}

func New(configPath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
