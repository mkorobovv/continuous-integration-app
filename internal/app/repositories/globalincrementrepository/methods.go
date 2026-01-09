package globalincrementrepository

import (
	"context"
	"time"
)

func (repo *GlobalIncrementRepository) Increment(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := repo.DB.Incr(ctx, key).Uint64()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
