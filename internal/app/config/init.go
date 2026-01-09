package config

import "github.com/ilyakaznacheev/cleanenv"

func New() (config Config, err error) {
	err = cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
