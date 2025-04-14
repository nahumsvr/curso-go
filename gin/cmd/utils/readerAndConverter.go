package utils

import (
	"encoding/json"
	"gin_http/cmd/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context) []byte {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el body"})
		return []byte{0}
	}
	return body
}

func ConvertToJson(c *gin.Context, body []byte, user *services.User) {
	err := json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
}
