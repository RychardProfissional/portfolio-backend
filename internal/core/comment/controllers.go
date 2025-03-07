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

	usecase := Usecase{}
	validate, err := usecase.ValidateCreate(&comment)
	if err != nil {
		log.Print("usecase.ValidateCreate: ", err.Error())
		c.JSON(http.StatusBadRequest,  gin.H{"error": err.Error()})
		return
	}

	created, err := usecase.Create(validate)
	if err != nil {
		log.Print("usecase.Create: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar registar no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, created)
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
