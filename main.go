package main

import (
	"log"

	"github.com/Ayyasy123/pt-aka-tech-vision/config"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.SetupUserRoutes(config.DB, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
