package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := Setup()
	r.Run()
}

func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	gr := r.Group("/api")
	RegisterProduct(gr)
	return r
}
