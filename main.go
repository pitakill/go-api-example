package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func basicGetRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola Mundo",
	})
}

func notFoundRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route not found",
	})
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")

	v1.GET("/", basicGetRoute)

	r.NoRoute(notFoundRoute)

	r.Run()
}
