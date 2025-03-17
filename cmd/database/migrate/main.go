package main

import (
	"github.com/RychardProfissional/portfolio-backend/config"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/internal/core/comment"
)

func init() {
	config.LoadEnv()
}

func main() {
	dbconn := db.Connect()

	dbconn.Debug().AutoMigrate(&comment.Repository{})
}