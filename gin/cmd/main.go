package main

import (
	"fmt"
	"gin_http/cmd/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.SetupRoutes(r)
	fmt.Println("Servidor escuchando en el puerto 3000")
	fmt.Println("http://localhost:3000")
	r.Run(":3000")
}
