package middlewares

import "github.com/gin-gonic/gin"

func GlobalMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
}
