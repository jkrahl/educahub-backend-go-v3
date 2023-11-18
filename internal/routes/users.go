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
		v1.POST("/me", RegisterUsernameHandler)
		v1.DELETE("/me", DeleteUserHandler)
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
	err = user.GetUserFromSub(sub)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if user.Username == "" {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func RegisterUsernameHandler(c *gin.Context) {
	type Body struct {
		Username string `json:"username"`
	}
	body := Body{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid body",
		})
		return
	}
	if body.Username == "" {
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

	user := models.User{}
	err = user.RegisterUser(sub, body.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func DeleteUserHandler(c *gin.Context) {
	sub, err := jwt.GetSubFromTokenFromRequest(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{}
	err = user.DeleteUser(sub)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
