package user

import (
	"gin_http/cmd/interfaces"
	"gin_http/cmd/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users []User

type User = interfaces.User

func GetAll(c *gin.Context) {
	c.JSON(200, users)
}

func Create(c *gin.Context) {
	body := utils.ReadBody(c)
	var user User
	utils.ConvertToJson(c, body, &user)
	user.ID = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	for i, user := range users {
		if user.ID == id {
			body := utils.ReadBody(c)
			utils.ConvertToJson(c, body, &user)
			users[i] = user
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}
