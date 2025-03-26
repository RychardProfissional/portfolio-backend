package config

import "github.com/gin-gonic/gin"

type Http struct{}

func (*Http) Gin() *gin.Engine {
	
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.RedirectTrailingSlash = true

	return router
}
