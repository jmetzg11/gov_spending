package main

import (
	"gov_spending/backend/database"
	"gov_spending/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":8080")
}
