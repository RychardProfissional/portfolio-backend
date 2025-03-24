package http

import (
	"log"

	"github.com/RychardProfissional/portfolio-backend/config"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/internal/core/comment"
	"github.com/RychardProfissional/portfolio-backend/internal/core/project"
)

type Routers struct{}

func (*Routers) Root() {
	configHttp := config.Http{}
	ginConf := configHttp.Gin()
	
	commentRouter := comment.Router{}
	projectRouter := project.Router{}
	
	commentRouter.Set(ginConf.Group("/comment"))
	projectRouter.Set(ginConf.Group("/project"))
	
	log.Print("Server in ", config.Env.SERVER.PORT)
	ginConf.Run("localhost:" + config.Env.SERVER.PORT)
}

func (*Routers) Boot() {
	config.LoadEnv()
	db.Connect()
}
