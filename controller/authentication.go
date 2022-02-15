package authentication_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	context.JSON(http.StatusCreated, "Success")
}

func SignIn(context *gin.Context) {
	context.JSON(http.StatusOK, "Success")
}
