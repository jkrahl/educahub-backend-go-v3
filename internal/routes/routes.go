package routes

import (
	"net/http"

	"educahub/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Educahub API",
		})
	})

	r.GET("/protected", authMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You are authorized",
		})
	})

	SetupUsersRoutes(r)
	SetupPostsRoutes(r)
}
