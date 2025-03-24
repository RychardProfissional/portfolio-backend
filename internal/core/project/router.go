package project

import "github.com/gin-gonic/gin"

type Router struct{}

func (*Router) Set(g *gin.RouterGroup)  {
	controller := Controller{}
	g.POST("/", controller.Create)
	g.GET("/by_id/:id", controller.GetByID)
	g.PATCH("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)
}