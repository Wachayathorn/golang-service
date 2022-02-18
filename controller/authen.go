package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database_connection "github.com/wachayathorn/golang-service/config/database-connection"
	dto "github.com/wachayathorn/golang-service/dto"
	models "github.com/wachayathorn/golang-service/model"
)

func SignIn(context *gin.Context) {
	var data dto.SignInDto

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := database_connection.PG_DB.Begin()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if err := tx.Model(&user).Where("username = ?", data.Username).Select(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Sign in failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sign in success"})
}
