package redis

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST"`
	User     string `yaml:"user" env:"USER"`
	Port     string `yaml:"port" env:"PORT"`
	Password string `yaml:"password" env:"PASSWORD"`
}

func New(config Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:            net.JoinHostPort(config.Host, config.Port),
		Password:        config.Password,
		Username:        config.User,
		DB:              0,
		ReadBufferSize:  1024 * 1024, // 1MiB
		WriteBufferSize: 1024 * 1024, // 1MiB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return rdb
}
