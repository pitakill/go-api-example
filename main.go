package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})

	router.Run()
}
