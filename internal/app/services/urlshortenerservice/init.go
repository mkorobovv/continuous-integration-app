package urlshortenerservice

import "context"

type URLShortenerService struct {
	shortenRepository         shortenRepository
	globalIncrementRepository globalIncrementRepository
}

type shortenRepository interface {
	Save(ctx context.Context, key, url string) (err error)
	Get(ctx context.Context, url string) (key string, err error)
}

type globalIncrementRepository interface {
	Increment(ctx context.Context, key string) (id int64, err error)
}

func New(shortenRepository shortenRepository, globalIncrementRepository globalIncrementRepository) *URLShortenerService {
	return &URLShortenerService{
		shortenRepository:         shortenRepository,
		globalIncrementRepository: globalIncrementRepository,
	}
}
