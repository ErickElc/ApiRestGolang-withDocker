package main

import (
	"ApiRestGolang/database"
	"ApiRestGolang/models"
	"ApiRestGolang/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor da api")
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1", Idade: 21},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2", Idade: 19},
	}
	database.ConnectDataBase()
	routes.HandleRequest()
}
