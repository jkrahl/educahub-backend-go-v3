package main

import (
	"log"

	"educahub/configs"
	"educahub/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Fatal(r.Run(":" + configs.GetViperString("PORT")))
}
