// main.go
package main

import (
	"golang-api/config"
	"golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDB()
	r := gin.Default()
	r.GET("/users", controllers.GetUsers)
	r.Run(":8080")
}
