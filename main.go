package main

import (
	"gov_spending/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":3000")
}
