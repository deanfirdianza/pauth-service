package main

import (
	"net/http"
	"os"
	"pauth/adapter"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/adapter", func(c *gin.Context) {
		adapter.Adapter()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"env":     os.Getenv("GIN_MODE"),
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
