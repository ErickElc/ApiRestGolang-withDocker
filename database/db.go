package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	stringConexao := "host=localhost user=root password=root dbname=root port=8231 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	} else {
		fmt.Println("Banco de dados conectado com sucesso")
	}
}
