package main

import (
	"github.com/adamwahyuh/coloroutfit/src/config"
	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/routes"
	"github.com/adamwahyuh/coloroutfit/src/seeders"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	route := gin.Default()

	database.ConnectDb()
	seeders.Seeder()

	routes.RegisterApiRoutes(route)

	host := config.GetEnv("APP_HOST")
	port := config.GetEnv("APP_PORT")

	route.Run(host + ":" + port)
}
