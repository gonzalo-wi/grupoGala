package main

import (
	"back/database"
	"back/routes"
	"log"
)

func main() {
	log.Println("Iniciando servidor...")

	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}
