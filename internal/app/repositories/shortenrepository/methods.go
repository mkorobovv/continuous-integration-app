package shortenrepository

import (
	"context"
	"fmt"
	"time"
)

func (repo *ShortenRepository) Save(ctx context.Context, key, url string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := repo.DB.Set(ctx, key, url, time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to save url: %w", err)
	}

	return nil
}

func (repo *ShortenRepository) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	url, err := repo.DB.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get url %s: %w", key, err)
	}

	return url, nil
}
