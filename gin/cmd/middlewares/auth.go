package middleware

import (
	"github.com/gin-gonic/gin"
)

func APIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API key is missing"})
			return
		}
		if apiKey != "ApiKey123456789" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API key"})
			return
		}
		c.Next()
	}

}
