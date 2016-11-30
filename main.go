package main

import (
	"encoding/json"
	"fmt"
	"os"

	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type Routes struct {
	Routes []Route
}

type Route struct {
	Name Name
}

type Name struct {
	Code    int
	Message string
}

func getMessagesRoutes() {
	file, _ := os.Open("message.json")
	decoder := json.NewDecoder(file)

	var route Route

	err := decoder.Decode(&route)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(route)
}

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
	getMessagesRoutes()

	r := gin.Default()

	v1 := r.Group("api/v1")

	v1.GET("/", basicGetRoute)

	r.NoRoute(notFoundRoute)

	r.Run()
}
