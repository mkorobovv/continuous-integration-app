package urlshortenerservice

import (
	"context"

	"github.com/mkorobovv/continuous-integration-app/internal/pkg/encoder"
)

func (svc *URLShortenerService) ShortenURL(ctx context.Context, url string) (string, error) {
	id, err := svc.globalIncrementRepository.Increment(ctx, "global_url_id")
	if err != nil {
		return "", err
	}

	key := encoder.Encode(id)

	err = svc.shortenRepository.Save(ctx, key, url)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (svc *URLShortenerService) GetURL(ctx context.Context, key string) (string, error) {
	url, err := svc.shortenRepository.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return url, nil
}
