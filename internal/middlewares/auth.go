package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikroongta8/choplinks/pkg/auth"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) <= len("Bearer ") {
		c.String(http.StatusUnauthorized, "Invalid Token!")
	}

	userClaims, err := auth.VerifyJWT(token[len("Bearer "):])
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid Token: "+err.Error())
	}
	c.Set("username", userClaims.Username)
	c.Next()
}
