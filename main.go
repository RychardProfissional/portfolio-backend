package main

import (
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/http"
)

var routers = http.Routers{}

func init() {
	routers.Boot()
}

func main() {
	routers.Root()	
}
