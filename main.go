package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	routes.GET("/", hello)

	routes.Run("localhost:3000")
}

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}
