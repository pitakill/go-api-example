package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Route not found",
		})
	})

	r.Run()
}
