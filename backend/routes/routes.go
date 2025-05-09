package routes

import (
	"gov_spending/backend/api"
	"gov_spending/backend/database"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	setupCORS(router)

	handler := api.NewHandler(database.DB)

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello from backend"})
		})
		apiRouter.GET("/foreign_aid", handler.GetForeignAid)
		apiRouter.GET("/agency", handler.GetAgencyData)
		apiRouter.GET("/function-spending", handler.GetFunctionSpending)
	}
}

func setupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow local frontend in dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func SetupStaticRoutes(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next() // Let API routes handle it
			return
		}

		// Try to serve static files
		filePath := "./frontend/build" + c.Request.URL.Path
		if _, err := os.Stat(filePath); err == nil {
			c.File(filePath)
			c.Abort()
			return
		}
		// If not a file, serve index.html (for SvelteKit routing)
		c.File("./frontend/build/index.html")
		c.Abort()
	})
}
