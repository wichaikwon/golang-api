// main.go
package main

import (
	"golang-api/config"
	"golang-api/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.ConnectToDB()
	r := gin.Default()
	r.GET("/users", controllers.GetUsers)
	r.Run(":8080")
}
