package db

import (
	"log"
	"sync"

	"github.com/RychardProfissional/portfolio-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var dbConn *gorm.DB

func Connect() (*gorm.DB) {
	once.Do(func() {
		configDb := config.Env.DB
		dsn := "host="+ configDb.HOST +
		" user="      + configDb.USER + 
		" password="  + configDb.PASS +
		" dbname="    + configDb.NAME +
		" port="      + configDb.PORT +
		" sslmode="   + configDb.SSLMODE

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("erro ao conectar com banco de dados")
		}
		dbConn = conn
	})
	return dbConn
}
