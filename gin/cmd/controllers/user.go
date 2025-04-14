package controllers

import (
	"gin_http/cmd/services"
	"gin_http/cmd/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (s *UserController) GetAll(c *gin.Context) {
	users := s.UserService.GetAll()
	c.JSON(200, users)
}

func (s *UserController) Create(c *gin.Context) {
	body := utils.ReadBody(c)
	var user services.User
	utils.ConvertToJson(c, body, &user)
	newUser := s.UserService.Create(user)

	c.JSON(http.StatusOK, newUser)
}

func (s UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	body := utils.ReadBody(c)
	user := services.User{}
	utils.ConvertToJson(c, body, &user)
	newUser, userErr := s.UserService.Update(id, user)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func (s *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	userErr := s.UserService.Delete(id)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
