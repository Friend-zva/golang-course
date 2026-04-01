package config

import (
	"time"

	env "github.com/Friend-zva/golang-course-task3/repo-stat/platform/env"
	grpcserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/grpcserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
)

type App struct {
	AppName string `yaml:"app_name" env:"APP_NAME" env-default:"repo-stat-subscriber"`
}

type Services struct {
	API string `yaml:"api" env:"API_ADDRESS" env-default:"localhost:8080"`
}

type Config struct {
	App      App               `yaml:"app"`
	Services Services          `yaml:"services"`
	GRPC     grpcserver.Config `yaml:"grpc"`
	Logger   logger.Config     `yaml:"logger"`
}

func MustLoad(path string) Config {
	var cfg Config
	env.MustLoad(path, &cfg)
	cfg.GRPC.Timeout = time.Duration(cfg.GRPC.TimeoutSec) * time.Second
	return cfg
}
