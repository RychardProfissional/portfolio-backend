package config

type ENV struct {
	DB     DB
	SERVER SERVER
}

type DB struct {
	HOST    string
	USER    string
	PASS    string
	NAME    string
	PORT    string
	SSLMODE string
}

type SERVER struct {
	PORT string
}