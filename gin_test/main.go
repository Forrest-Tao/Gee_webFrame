package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusNotFound, "NOT FOUND ")
	})

	engine.Run()
}
