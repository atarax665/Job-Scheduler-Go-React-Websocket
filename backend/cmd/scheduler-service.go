package main

import (
	"log"
	"scheduler-service/internal/controller/jobs"
	"scheduler-service/internal/controller/status"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	log.Printf("Starting Job Scheduling server")
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(corsMiddleware())

	jobs.InitRoutes(router)
	status.InitRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
