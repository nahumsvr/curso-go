package main

import (
	"fmt"
	middleware "gin_http/cmd/middlewares"
	"gin_http/cmd/routes"
	"gin_http/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	// services
	userService := services.NewUserService()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.SetupRoutes(r, userService)
	fmt.Println("Servidor escuchando en el puerto 3000")
	fmt.Println("http://localhost:3000")
	r.Run(":3000")
}
