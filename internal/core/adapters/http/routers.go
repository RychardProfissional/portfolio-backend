package http

import (
	"github.com/RychardProfissional/portfolio-backend/config"
	"github.com/RychardProfissional/portfolio-backend/internal/core/comment"
)

type Routers struct{}

func (*Routers) Root() {
	configHttp := config.Http{}
	ginConf := configHttp.Gin()
	
	commentRouter := comment.Router{}
	
	commentRouter.Set(ginConf.Group("/comment"))
	
	ginConf.Run("localhost:" + config.Env.SERVER.PORT)
}

func (*Routers) Boot() {
	config.LoadEnv()
}
