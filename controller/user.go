package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	models "github.com/wachayathorn/golang-service/model"

	dto "github.com/wachayathorn/golang-service/dto"

	database_connection "github.com/wachayathorn/golang-service/config/database-connection"
)

func CreateUser(context *gin.Context) {
	var data dto.CreateUserDto

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database_connection.GORM_DB.Begin()

	user := models.User{}
	if err := database_connection.GORM_DB.Where("username = ?", data.Username).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist!"})
		return
	}

	user = models.User{Id: strconv.FormatInt(time.Now().Unix(), 10), FirstName: data.FirstName, LastName: data.LastName, Username: data.Username}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	context.JSON(http.StatusCreated, user)
}

func GetUser(context *gin.Context) {
	userList := []models.User{}
	database_connection.GORM_DB.Find(&userList)
	context.JSON(http.StatusOK, userList)
}

func GetUserById(context *gin.Context) {
	id := context.Param("id")

	user := models.User{}
	if err := database_connection.GORM_DB.Where("id = ?", id).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	id := context.Param("id")

	user := models.User{}
	if err := database_connection.GORM_DB.Where("id = ?", id).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database_connection.GORM_DB.Begin()

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusOK, user)
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	user := models.User{}

	tx := database_connection.GORM_DB.Begin()

	if err := database_connection.GORM_DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	tx.Commit()

	context.JSON(http.StatusOK, gin.H{"delete": "success"})
}
