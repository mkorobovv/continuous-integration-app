package globalincrementrepository

import "github.com/redis/go-redis/v9"

type GlobalIncrementRepository struct {
	DB *redis.Client
}

func New(db *redis.Client) *GlobalIncrementRepository {
	return &GlobalIncrementRepository{
		DB: db,
	}
}
