package routes

import (
	"gin_http/cmd/controllers"
	middleware "gin_http/cmd/middlewares"
	"gin_http/cmd/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userService *services.UserService) {
	admin := r.Group("/admin")
	admin.Use(middleware.APIKey())

	userController := controllers.NewUserController(userService)

	admin.GET("/users", userController.GetAll)
	admin.POST("/users", userController.Create)
	admin.PUT("/users/:id", userController.Update)
	admin.DELETE("/users/:id", userController.Delete)
}
