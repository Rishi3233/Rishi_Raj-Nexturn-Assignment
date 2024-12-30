package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidateProduct checks if the required fields are present in the product request
func ValidateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product map[string]interface{}
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}
		if product["name"] == nil || product["price"] == nil || product["stock"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			c.Abort()
			return
		}
		c.Next()
	}
}
