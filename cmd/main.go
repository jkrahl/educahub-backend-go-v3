package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jkrahl/educahub-api/internal/routes"
	"github.com/spf13/viper"
)

func main() {
	// Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	r := gin.Default()

	routes.SetupRoutes(r)
	log.Fatal(r.Run(":" + viper.GetString("port")))
}
