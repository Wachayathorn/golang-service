package main

import (
	"github.com/gin-gonic/gin"

	authentication_controller "github.com/wachayathorn/golang-service/controller"
)

func main() {
	r := gin.Default()

	r.GET("/signup", authentication_controller.SignUp)
	r.GET("/signin", authentication_controller.SignIn)

	r.Run()
}
