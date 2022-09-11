package main

import (
	"jwt-gin/models"
	"jwt-gin/routes"
)

func main() {
	models.ConnectDataBase()
	routes.Handler()
}
