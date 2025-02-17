package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Inside goauth packages main file")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", home)
	router.NoRoute(func(c *gin.Context) {
		fmt.Println("Unfortunately page not found")
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func home(c *gin.Context) {
	fmt.Println("At home page")             // This prints in terminal
	c.String(http.StatusOK, "At home page") // This prints in webpage
}
