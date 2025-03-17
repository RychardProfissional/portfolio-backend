package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	log.Print("load")
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		log.Print(filename)
		envFile := filepath.Dir(filename) + "\\..\\..\\..\\..\\.env"
		envFile = strings.ReplaceAll(envFile, "\\", "/")
		log.Print(envFile)
		// if _, err := os.Stat(envFile); os.IsNotExist(err) {
			godotenv.Load(envFile)
		// }
	}

	Env = ENV{}
	Env.DB = db{
		HOST: os.Getenv("DB_HOST"),
		USER: os.Getenv("DB_USER"),
		PASS: os.Getenv("DB_PASS"),
		NAME: os.Getenv("DB_NAME"),
		PORT: os.Getenv("DB_PORT"),
		SSLMODE: os.Getenv("DB_SSLMODE"),
	}

	Env.SERVER = server{}
	Env.SERVER.PORT = os.Getenv("SYS_PORT")
	if Env.SERVER.PORT == "" {
		Env.SERVER.PORT = "8080"
	} 
}