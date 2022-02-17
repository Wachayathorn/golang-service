package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gonanoid "github.com/matoous/go-nanoid/v2"

	models "github.com/wachayathorn/golang-service/model"
)

var userList []models.User

func CreateUser(context *gin.Context) {
	var data models.User

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := gonanoid.New()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Id = string(id)

	userList = append(userList, data)

	context.JSON(http.StatusCreated, data)
}

func GetUser(context *gin.Context) {
	context.JSON(http.StatusOK, userList)
}

func GetUserById(context *gin.Context) {
	id := context.Param("id")

	for _, user := range userList {
		if user.Id == id {
			context.JSON(http.StatusOK, user)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func UpdateUser(context *gin.Context) {
	id := context.Param("id")

	for index, user := range userList {
		if user.Id == id {
			if err := context.ShouldBindJSON(&user); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			userList[index] = user
			context.JSON(http.StatusOK, user)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	for index, user := range userList {
		if user.Id == id {
			userList = append(userList[:index], userList[index+1:]...)
			context.JSON(http.StatusOK, gin.H{"deleted": true})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}
