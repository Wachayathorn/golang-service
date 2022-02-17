package main

import (
	"github.com/gin-gonic/gin"

	user_controller "github.com/wachayathorn/golang-service/controller"

	database_connection "github.com/wachayathorn/golang-service/config"
)

func main() {
	// Connect Database
	database_connection.ConnectDatabase()

	// Initialize Router
	r := gin.Default()

	r.POST("/user", user_controller.CreateUser)
	r.GET("/user", user_controller.GetUser)
	r.GET("/user/:id", user_controller.GetUserById)
	r.PUT("/user/:id", user_controller.UpdateUser)
	r.DELETE("/user/:id", user_controller.DeleteUser)

	r.Run()
}
