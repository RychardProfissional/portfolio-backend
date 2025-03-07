package config

type ENV struct {
	DB     db
	SERVER server
}

type db struct {
	HOST    string
	USER    string
	PASS    string
	NAME    string
	PORT    string
	SSLMODE string
}

type server struct {
	PORT string
}