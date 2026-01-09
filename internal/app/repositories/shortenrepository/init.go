package shortenrepository

import "github.com/redis/go-redis/v9"

type ShortenRepository struct {
	DB *redis.Client
}

func New(db *redis.Client) *ShortenRepository {
	return &ShortenRepository{
		DB: db,
	}
}
