package main

import (
	"github.com/Nasaee/go-gin-rest-api/db"
	"github.com/Nasaee/go-gin-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
