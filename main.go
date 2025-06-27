package main

import (
	"github.com/FedeGonzalez84/go-rest-api/db"
	"github.com/FedeGonzalez84/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
