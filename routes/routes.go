package routes

import (
	"ZumaAkbarID/backend-api/controllers"
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

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	return router
}
