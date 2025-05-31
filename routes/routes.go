package routes

import (
	"ZumaAkbarID/backend-api/controllers"
	"ZumaAkbarID/backend-api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Hello World!",
		})
	})

	// Auth
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	// Users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)

	return router
}
