package config

import "github.com/mkorobovv/continuous-integration-app/internal/app/infrastructure/redis"

type Config struct {
	App       App       `yaml:"app"`
	Databases Databases `yaml:"databases"`
	HTTP      HTTP      `yaml:"http"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Databases struct {
	Shorten redis.Config `yaml:"shorten" env-prefix:"SHORTEN_"`
}

type HTTP struct {
	Port string `yaml:"port"`
}
