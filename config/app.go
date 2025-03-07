package config

import "os"

func LoadEnv() {
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