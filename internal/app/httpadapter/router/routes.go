package router

import (
	"github.com/mkorobovv/continuous-integration-app/internal/app/httpadapter/handlers"
)

const apiV1Prefix = "/api/v1"

func (r *Router) SetupRoutes(handler *handlers.Handler) {
	r.engine.GET("/:key", handler.Redirect)

	v1 := r.engine.Group(apiV1Prefix)
	{
		v1.PUT("/shorten", handler.ShortenURL)
	}
}
