package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database_connection "github.com/wachayathorn/golang-service/config/database-connection"
	dto "github.com/wachayathorn/golang-service/dto"
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

	if user := tx.QueryRow("SELECT * FROM users WHERE username = ?", data.Username); user == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Sign in failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Sign in success"})
}
