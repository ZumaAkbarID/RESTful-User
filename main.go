package main

import (
	"ZumaAkbarID/backend-api/config"
	"ZumaAkbarID/backend-api/database"
	"ZumaAkbarID/backend-api/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
