package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()                                      // Record the start time
		method := c.Request.Method                                   // Get the HTTP method
		path := c.Request.URL.Path                                   // Get the request path
		clientIP := c.ClientIP()                                     // Get the client IP address
		log.Printf("Request: %s %s from %s", method, path, clientIP) // Log the request details
		c.Next()                                                     // Call the next handler in the chain
		endTime := time.Now()                                        // Record the end time
		duration := endTime.Sub(startTime)                           // Calculate the duration
		statusCode := c.Writer.Status()                              // Get the response status code
		log.Printf("Response: %d in %v", statusCode, duration)       // Log the response details
	}

}
