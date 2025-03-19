package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	godotenv.Load()

	Env = ENV{}
	Env.DB = DB{
		HOST: os.Getenv("DB_HOST"),
		USER: os.Getenv("DB_USER"),
		PASS: os.Getenv("DB_PASS"),
		NAME: os.Getenv("DB_NAME"),
		PORT: os.Getenv("DB_PORT"),
		SSLMODE: os.Getenv("DB_SSLMODE"),
	}

	Env.SERVER = SERVER{}
	Env.SERVER.PORT = os.Getenv("SYS_PORT")
	if Env.SERVER.PORT == "" {
		Env.SERVER.PORT = "8080"
	} 
}