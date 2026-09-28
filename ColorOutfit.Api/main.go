package main

import (
	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/routes"
	"github.com/adamwahyuh/coloroutfit/src/seeders"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	routes.RegisterApiRoutes(route)

	database.ConnectDb()
	seeders.Seeder()
	route.Run()
}
