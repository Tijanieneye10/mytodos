package main

import (
	"github.com/Tijanieneye10/database"
	"github.com/Tijanieneye10/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//Create tables
	database.InitDB()

	//Start http server
	server := gin.Default()

	//Register route
	routes.RegisterRoutes(server)

	err := server.Run()
	if err != nil {
		return
	}

}
