package main

import (
	"github.com/BrunoIstvan/api-rest-go-gin/database"
	"github.com/BrunoIstvan/api-rest-go-gin/routes"
)

func main() {

	database.ConnectDatabase()
	routes.HandleRequests()

}
