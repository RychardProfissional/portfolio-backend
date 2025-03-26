package main

import (
	"log"
	"os"

	"github.com/RychardProfissional/portfolio-backend/config"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/models"
	"github.com/lpernett/godotenv"
)

func init() {
	LoadEnv()
}

func main() {
	dbconn := db.Connect()

	dbconn.Debug().AutoMigrate(&models.Project{})
	dbconn.Debug().AutoMigrate(&models.Comment{})
}

func LoadEnv() {
	log.Print("load")
	envFile :=  "./../../../.env"
	godotenv.Load(envFile)

	config.Env = config.ENV{}
	config.Env.DB = config.DB{
		HOST: os.Getenv("DB_HOST"),
		USER: os.Getenv("DB_USER"),
		PASS: os.Getenv("DB_PASS"),
		NAME: os.Getenv("DB_NAME"),
		PORT: os.Getenv("DB_PORT"),
		SSLMODE: os.Getenv("DB_SSLMODE"),
	}

	config.Env.SERVER = config.SERVER{}
	config.Env.SERVER.PORT = os.Getenv("SYS_PORT")
	if config.Env.SERVER.PORT == "" {
		config.Env.SERVER.PORT = "8080"
	} 
}