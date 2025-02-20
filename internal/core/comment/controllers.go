package comment

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (*Controller) Create(c *gin.Context) {
	var comment Entitie

	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Print("c.ShouldBindJSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parametros invalidos"})
		return
	}

	
}

func (*Controller) GetByID(c *gin.Context) {
	// TODO:
} 

func (*Controller) GetByProjectID(c *gin.Context) {
	// TODO:
} 

func (*Controller) UpdateText(c *gin.Context) {
	// TODO:
} 

func (*Controller) Delete(c *gin.Context) {
	// TODO:
} 
