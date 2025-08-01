package controllers

import (
	"back/database"
	"back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterIP(c *gin.Context) {
	ip := c.ClientIP()
	visitor := models.VisitorIP{IP: ip}
	database.DB.Create(&visitor)
	c.JSON(http.StatusOK, gin.H{"status": "ip registrada", "ip": ip})
}

// Nuevo endpoint para obtener todas las IPs
func GetIPs(c *gin.Context) {
	var ips []models.VisitorIP
	database.DB.Order("created_at desc").Find(&ips)
	c.JSON(http.StatusOK, ips)
}
