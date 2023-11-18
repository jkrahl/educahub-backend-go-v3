package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkrahl/educahub-api/internal/middleware"
	"github.com/spf13/viper"
)

func SetupRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": viper.GetString("default-message"),
		})
	})

	r.GET("/protected", authMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You are authorized",
		})
	})

	SetupUsersRoutes(r)
}
