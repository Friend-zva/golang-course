package config

import (
	"time"

	env "github.com/Friend-zva/golang-course-task3/repo-stat/platform/env"
	httpserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/httpserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
)

type App struct {
	AppName string `yaml:"app_name" env:"APP_NAME" env-default:"repo-stat-api"`
}

type Services struct {
	Subscriber string `yaml:"subscriber" env:"SUBSCRIBER_ADDRESS" env-default:"localhost:8081"`
	Processor  string `yaml:"processor" env:"PROCESSOR_ADDRESS" env-default:"localhost:8082"`
}

type Config struct {
	App      App               `yaml:"app"`
	Services Services          `yaml:"services"`
	HTTP     httpserver.Config `yaml:"http"`
	Logger   logger.Config     `yaml:"logger"`
}

func MustLoad(path string) Config {
	var cfg Config
	env.MustLoad(path, &cfg)
	cfg.HTTP.Timeout = time.Duration(cfg.HTTP.TimeoutSec) * time.Second
	return cfg
}
