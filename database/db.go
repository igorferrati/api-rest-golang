package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectaBnacoDB() {
	stringConnection := "host=localhost user=root password=root dbname=root port=54321 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados, err: ", err)
	}
}
