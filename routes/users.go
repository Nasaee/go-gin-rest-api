package routes

import (
	"log"
	"net/http"

	"github.com/Nasaee/go-gin-rest-api/models"
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
