package main

import (
	"github.com/eze-echu/guitapp/config"
	//"github.com/eze-echu/guitapp/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file, Using defaults")
	}
}
func main() {
	Conf := config.New()
	//db.InitDB(Conf.DB)
	router := gin.Default()
	//routes.RegisterRoutes(router)

	err := router.Run(":5000")
	if err != nil {
		return
	}
}
