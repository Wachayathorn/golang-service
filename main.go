package main

import (
	"github.com/gin-gonic/gin"

	database_connection "github.com/wachayathorn/golang-service/config/database-connection"
	controller "github.com/wachayathorn/golang-service/controller"
)

func main() {
	// Connect Database
	go database_connection.ConnectDatabaseByGORM()
	go database_connection.ConnectDatabaseByPQ()

	// Initialize Router
	r := gin.Default()

	// User Router
	r.POST("/user", controller.CreateUser)
	r.GET("/user", controller.GetUser)
	r.GET("/user/:id", controller.GetUserById)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	// Authentication Router
	r.POST("/signin", controller.SignIn)

	r.Run()
}
