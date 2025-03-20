package routes

import (
	"gin_http/cmd/controller/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", user.GetAll)
	r.POST("/users", user.Create)
	r.PUT("/users/:id", user.Update)
	r.DELETE("/users/:id", user.Delete)
}
