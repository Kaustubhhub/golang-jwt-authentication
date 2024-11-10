package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api-1", func(c *gin.Context) {
		fmt.Println("It reached to api-1")
		c.JSON(200, gin.H{"success": "Access granted for api-1."})
	})

	router.GET("/api-2", func(c *gin.Context) {
		fmt.Println("It reached to api-2")
		c.JSON(200, gin.H{"success": "Access granted for api-2."})
	})

	router.Run(":" + port)
}
