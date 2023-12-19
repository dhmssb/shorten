package router

import (
	handlers "amartha/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/shorten", handlers.ShortenURL)
	router.PUT("/slug/:slug", handlers.ModifySlug)
	router.GET("/:slug", handlers.RedirectToURL)
	router.GET("/urls", handlers.GetShortURLs)

	return router
}
