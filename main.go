package main

import (
	"log"
	"net/http"

	"github.com/Nasaee/go-gin-rest-api/db"
	"github.com/Nasaee/go-gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events. Try again later."})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	// ShouldBindJSON : มันมีหน้าที่ อ่าน JSON จาก request body แล้วแปลงเป็น struct ที่คุณส่งเข้าไป
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}
