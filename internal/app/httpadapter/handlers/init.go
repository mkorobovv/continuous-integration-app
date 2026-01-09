package handlers

import "context"

type Handler struct {
	urlShortenerService urlShortenerService
}

type urlShortenerService interface {
	ShortenURL(ctx context.Context, url string) (key string, err error)
	GetURL(ctx context.Context, key string) (url string, err error)
}

func New(urlShortenerService urlShortenerService) *Handler {
	return &Handler{
		urlShortenerService: urlShortenerService,
	}
}
