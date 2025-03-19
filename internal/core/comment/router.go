package comment

import "github.com/gin-gonic/gin"

type Router struct{}

func (*Router) Set(g *gin.RouterGroup)  {
	controller := Controller{}
	g.POST("/", controller.Create)
	g.GET("/by_id/:id", controller.GetByID)
	// g.GET("/by_project_id/:project_id", controller.GetByProjectID)
	g.PATCH("/:id/:ip", controller.UpdateText)
	g.DELETE("/:id", controller.Delete)
}