package comment

import (
	"log"
	"net/http"

	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct{}

func (*Controller) Create(c *gin.Context) {
	var comment entities.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Print("c.ShouldBindJSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "parametros invalidos"})
		return
	}

	usecase := Usecase{}
	validated, err := usecase.ValidateCreate(&comment)
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
	commentUUID, err := uuid.Parse(id)
	if commentUUID == uuid.Nil || err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	usecase := Usecase{}
	comment, err := usecase.GetByID(id)
	if err != nil {
		log.Print("usecase.GetByID: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao obter registros no banco de dados"})
		return
	}

	if comment == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, comment)
} 

func (*Controller) GetByProjectID(c *gin.Context) {
	// TODO
} 

func (*Controller) UpdateText(c *gin.Context) {
	id := c.Param("id")
	ip := c.Param("ip")

	text := c.Query("text")

	usecase := Usecase{}
	Validated, err := usecase.ValidateUpdateText(id, ip, text)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = usecase.UpdateText(id, &text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno na API"})
		return
	}

	Validated.Text = text
	c.JSON(http.StatusOK, Validated)
} 

func (*Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	commentUUID, err := uuid.Parse(id)
	if commentUUID == uuid.Nil || err != nil  {
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
