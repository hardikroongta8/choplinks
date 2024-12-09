package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	res := response{Error: message}
	log.Println("ERROR:", message+",", "STATUS:", statusCode)
	c.JSON(statusCode, res)
}

func SendSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	res := response{Data: data}
	c.JSON(statusCode, res)
}
