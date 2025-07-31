package controllers

import (
	"back/database"
	"back/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ReportVictim(c *gin.Context) {
	var victim models.Victim

	if err := c.ShouldBindJSON(&victim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	victim.IPAddress = c.ClientIP()
	victim.CreatedAt = time.Now()

	database.DB.Create(&victim)
	c.JSON(http.StatusOK, gin.H{"status": "registrado"})
}

// Endpoint para listar pedidos
func GetPedidos(c *gin.Context) {
	var pedidos []models.Pedido
	database.DB.Find(&pedidos)
	c.JSON(http.StatusOK, pedidos)
}

func GetVictims(c *gin.Context) {
	var victims []models.Victim
	database.DB.Find(&victims)
	c.JSON(http.StatusOK, victims)
}

// Nuevo endpoint para guardar pedidos
func CreatePedido(c *gin.Context) {
	var pedido models.Pedido

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pedido.Fecha = time.Now()

	database.DB.Create(&pedido)
	c.JSON(http.StatusOK, gin.H{"status": "pedido registrado"})
}
