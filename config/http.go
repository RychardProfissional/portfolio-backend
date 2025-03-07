package config

import "github.com/gin-gonic/gin"

type Http struct{}

func (*Http) Gin() *gin.Engine {
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.RedirectTrailingSlash = true

	return router
}
