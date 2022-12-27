package main

import (
	"API_GO_REST/database"
	"API_GO_REST/models"
	"API_GO_REST/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Olyver", Historia: "Vocalista"},
		{Id: 2, Nome: "Olyver", Historia: "Vocalista"},
	}
	database.ConectaComBancoDeDados()

	fmt.Println("Inicinado o servidor rest com go...")

	routes.HandleResquest()
}
