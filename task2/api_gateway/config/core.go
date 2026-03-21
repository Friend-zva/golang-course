package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env             string           `yaml:"env" env-default:"local"`
	HTTPServer      HTTPServerConfig `yaml:"http_server"`
	CollectorClient CollectorConfig  `yaml:"collector_client"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type CollectorConfig struct {
	Address string `yaml:"address" env-default:"localhost:8081"`
}

func MustLoad(pathConfig string) *Config {
	if _, err := os.Stat(pathConfig); os.IsNotExist(err) {
		log.Fatalf("Failed to find config file (%s)", pathConfig)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(pathConfig, &cfg); err != nil {
		log.Fatalf("Failed to read config (%s)", err)
	}

	return &cfg
}
