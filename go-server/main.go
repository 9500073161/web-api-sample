package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go")
	})

	router.GET("/prashob", func(c *gin.Context) {
		c.String(200, "Hello Prashob %s", os.Getenv("NAME"))
	})

	router.GET("/kanchi", func(c *gin.Context) {
		c.String(200, "Hello Kanchi %s", os.Getenv("NAME"))
	})

	router.Run(":" + port)
}
