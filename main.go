package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"shorter/db"
	"shorter/routes"
)

func main() {
	db.IntiDB()

	server := gin.Default()

	routes.Routes(server)

	log.Println("Server running on :1010")
	server.Run(":1010")
}
