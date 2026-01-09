package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func New() *Router {
	return &Router{
		engine: gin.New(),
	}
}

func (r *Router) Router() http.Handler {
	return r.engine.Handler()
}
