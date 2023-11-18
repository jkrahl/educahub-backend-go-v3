package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkrahl/educahub-api/internal/jwt"
	"github.com/jkrahl/educahub-api/internal/middleware"
	"github.com/jkrahl/educahub-api/internal/models"
)

func SetupUsersRoutes(r *gin.Engine) {
	jwtMiddleware, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/auth", jwtMiddleware)
	{
		v1.GET("/me", GetMyUsernameHandler)
		v1.POST("/me", RegisterOrUpdateUsernameHandler)
	}
}

func GetMyUsernameHandler(c *gin.Context) {
	sub, err := jwt.GetSubFromTokenFromRequest(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	models.DB.Where("sub = ?", sub).First(&user)
	if user.Sub == "" {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func RegisterOrUpdateUsernameHandler(c *gin.Context) {
	usernameParam := c.PostForm("username")
	if usernameParam == "" {
		c.JSON(400, gin.H{
			"error": "username is required",
		})
		return
	}
	sub, err := jwt.GetSubFromTokenFromRequest(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	existentUser := models.User{}
	models.DB.Where("sub = ?", sub).First(&existentUser)
	if existentUser.Sub != "" {
		c.JSON(400, gin.H{
			"error": "user already exists",
		})
		return
	}

	user := models.User{}
	user.Sub = sub
	user.Username = usernameParam
	models.DB.Create(&user)

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}
