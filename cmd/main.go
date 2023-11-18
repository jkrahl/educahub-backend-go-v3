package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jkrahl/educahub-api/configs"
	"github.com/jkrahl/educahub-api/internal/models"
	"github.com/jkrahl/educahub-api/internal/routes"
	"github.com/spf13/viper"
)

func main() {
	configs.SetupViper("../configs")

	models.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Fatal(r.Run(":" + viper.GetString("port")))
}
