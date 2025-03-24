package project

import (
	"log"
	"net/http"

	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct{}

func (*Controller) Create(c *gin.Context) {
	var project entities.Project
	
	if err := c.ShouldBindJSON(&project); err != nil {
		log.Print("c.ShouldBindJSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parametros invalidos"})
		return
	}

	usecase := Usecase{}

	validated, err := usecase.ValidateCreate(&project); 
	if err != nil {
		log.Print("usecase.ValidateCreate: ", err.Error())
		c.JSON(http.StatusBadRequest,  gin.H{"error": err.Error()})
		return
	}

	created, err := usecase.Create(validated)
	if err != nil {
		log.Print("usecase.Create: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar registar no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, created)
}


func (*Controller) GetByID(c *gin.Context) {
	id := c.Param("id")
	projectUUID, err := uuid.Parse(id)
	if projectUUID == uuid.Nil || err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	usecase := Usecase{}
	project, err := usecase.GetByID(id)
	if err != nil {
		log.Print("usecase.GetByID: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao obter registros no banco de dados"})
		return
	}

	if project == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, project)
} 

func (*Controller) Update(c *gin.Context) {
	id := c.Param("id")
	projectUUID, err := uuid.Parse(id)
	if projectUUID == uuid.Nil || err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	var project entities.Project
	
	if err := c.ShouldBindJSON(&project); err != nil {
		log.Print("c.ShouldBindJSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parametros invalidos"})
		return
	}

	usecase := Usecase{}

	validated, err := usecase.ValidateUpdate(id, &project); 
	if err != nil {
		log.Print("usecase.ValidateUpdate: ", err.Error())
		c.JSON(http.StatusBadRequest,  gin.H{"error": err.Error()})
		return
	}

	updated, err := usecase.Update(validated)
	if err != nil {
		log.Print("usecase.Update: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar registar no banco de dados"})
		return
	}

	c.JSON(http.StatusOK, updated)
} 

func (*Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	projectUUID, err := uuid.Parse(id)
	if projectUUID == uuid.Nil || err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	usecase := Usecase{}
	err = usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno na API"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
} 