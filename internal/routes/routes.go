package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkrahl/educahub-api/internal/middleware"
	"github.com/jkrahl/educahub-api/internal/models"
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

	r.GET("/checkIfDBConnected", func(c *gin.Context) {
		if models.GetDBInstance() == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "DB is not connected",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "DB is connected",
			})
		}
	})

	SetupUsersRoutes(r)
}
