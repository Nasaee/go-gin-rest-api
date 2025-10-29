package routes

import (
	"net/http"
	"strconv"

	"github.com/Nasaee/go-gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent (context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	
	event, err :=models.GetEventById(eventId)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event."})
		return
	}
}

func cancelRegistration (context *gin.Context) {}