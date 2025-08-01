package routes

import (
	"back/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.0.55:8094"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.POST("/report", controllers.ReportVictim)
		api.GET("/results", controllers.GetVictims)
		api.GET("/pedidos", controllers.GetPedidos)
		api.POST("/pedido", controllers.CreatePedido)
		api.POST("/register-ip", controllers.RegisterIP)
		api.GET("/ips", controllers.GetIPs)
	}

	return r
}
