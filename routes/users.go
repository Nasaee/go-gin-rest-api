package routes

import (
	"log"
	"net/http"

	"github.com/Nasaee/go-gin-rest-api/models"
	"github.com/Nasaee/go-gin-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	// ShouldBindJSON : มันมีหน้าที่ อ่าน JSON จาก request body แล้วแปลงเป็น struct ที่คุณส่งเข้าไป
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		log.Println(err)
		// context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticated user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
