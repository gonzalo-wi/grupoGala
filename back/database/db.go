package database

import (
	"back/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("pishing.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar la base de datos:", err)
	}
	DB.AutoMigrate(&models.Pedido{})
	DB.AutoMigrate(&models.Victim{})
	DB.AutoMigrate(&models.VisitorIP{})
	log.Println("Conexión a la base de datos establecida y modelo migrado")
}
