package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Logging middleware to log all incoming requests
func LogRequest(c *gin.Context) {
	fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL)
	c.Next()
}
