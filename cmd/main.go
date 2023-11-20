package main

import (
	"log"

	"educahub/configs"
	"educahub/internal/models"
	"educahub/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs.SetupViper("./configs")

	models.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Fatal(r.Run(":" + viper.GetString("port")))
}
