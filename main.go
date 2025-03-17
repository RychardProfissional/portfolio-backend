package main

import (
	"log"

	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/http"
)

var routers = http.Routers{}

func init() {
	log.Print("init")
	routers.Boot()
}

func main() {
	routers.Root()	
}
