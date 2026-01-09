package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenURL(ctx *gin.Context) {
	var dtoIn ShortenRequest

	err := ctx.ShouldBindQuery(&dtoIn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	shorten, err := h.urlShortenerService.ShortenURL(ctx, dtoIn.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(
		http.StatusOK,
		ShortenResponse{
			URL: shorten,
		},
	)
}

type ShortenRequest struct {
	URL string `form:"url" binding:"required,url"`
}

type ShortenResponse struct {
	URL string `json:"url"`
}

func (h *Handler) Redirect(ctx *gin.Context) {
	key := ctx.Param("key")

	url, err := h.urlShortenerService.GetURL(ctx, key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
