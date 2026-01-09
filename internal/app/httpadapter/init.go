package httpadapter

import (
	"context"
	"net/http"
	"time"

	"github.com/mkorobovv/continuous-integration-app/internal/app/httpadapter/handlers"
	"github.com/mkorobovv/continuous-integration-app/internal/app/httpadapter/router"
	"github.com/mkorobovv/continuous-integration-app/internal/app/services/urlshortenerservice"
	"golang.org/x/sync/errgroup"
)

type HttpAdapter struct {
	server *http.Server
}

func New(urlShortenerService *urlshortenerservice.URLShortenerService) *HttpAdapter {
	handler := handlers.New(urlShortenerService)

	r := router.New()

	r.SetupRoutes(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Router(),
	}

	return &HttpAdapter{
		server: server,
	}
}

func (r *HttpAdapter) Start(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return r.server.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return r.server.Shutdown(ctx)
	})

	return g.Wait()
}
