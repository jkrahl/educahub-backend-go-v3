package routes

import (
	"educahub/internal/jwt"
	"educahub/internal/middleware"
	"educahub/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(r *gin.Engine) {
	AuthRequired, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/users", AuthRequired)
	{
		v1.GET("/", middleware.CheckIfUserExists(), GetMyUsernameHandler)
		v1.POST("/", RegisterUsernameHandler)
		v1.DELETE("/", middleware.CheckIfUserExists(), DeleteUserHandler)
	}
}

func GetMyUsernameHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func RegisterUsernameHandler(c *gin.Context) {
	type Body struct {
		Username string `json:"username" binding:"required"`
	}
	body := Body{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid body",
		})
		return
	}

	sub, err := jwt.GetSubFromTokenFromContext(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Username: body.Username,
		Sub:      sub,
	}
	err = user.Create()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func DeleteUserHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	err := user.Delete()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
