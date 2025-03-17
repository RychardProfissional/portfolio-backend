package http

import (
	"log"

	"github.com/RychardProfissional/portfolio-backend/config"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/internal/core/comment"
)

type Routers struct{}

func (*Routers) Root() {
	configHttp := config.Http{}
	ginConf := configHttp.Gin()
	
	commentRouter := comment.Router{}
	
	commentRouter.Set(ginConf.Group("/comment"))
	
	log.Print("Server in ", config.Env.SERVER.PORT)
	ginConf.Run("localhost:" + config.Env.SERVER.PORT)
}

func (*Routers) Boot() {
	config.LoadEnv()
	db.Connect()
}
