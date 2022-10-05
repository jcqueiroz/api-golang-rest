package main

import (
	"ALURA/jcqueiroz2/api-go-rest/database"
	"ALURA/jcqueiroz2/api-go-rest/models"
	"ALURA/jcqueiroz2/api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	database.ConnectWithDataBase()
	fmt.Println("Start server Rest with Golang ")
	routes.HandleRequest()

}
