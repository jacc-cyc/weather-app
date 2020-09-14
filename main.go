package main

import (
	"gin-weather-app/weather"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/", weather.Home)
	router.GET("/info/", weather.Info)
	router.GET("/info/:date", weather.DayInfo)
	router.GET("/warninginfo/", weather.WarningInfo)

	router.Run(":" + port)
}
