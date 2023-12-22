package main

import (
	"fmt"

	"github.com/igorferrati/api-rest-golang/models"
	"github.com/igorferrati/api-rest-golang/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando projeto!")
	routes.HandleResquest()
}
