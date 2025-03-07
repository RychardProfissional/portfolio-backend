package main

import "github.com/RychardProfissional/portfolio-backend/internal/core/adapters/http"

func main() {
	routers := http.Routers{}
	routers.Boot()
	routers.Root()	
}
