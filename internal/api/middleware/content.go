package middleware

import "github.com/gin-gonic/gin"

func ContentType(c *gin.Context) {
	c.Header("Content-Type", "application/json")
}
