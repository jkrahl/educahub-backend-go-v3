package middleware

import (
	jwtutils "educahub/internal/jwt"
	"educahub/internal/models"
	"log"

	gin "github.com/gin-gonic/gin"
)

// This function is used to check if a user exists in the database.
// The request requires an Authorization header with a valid JWT.
// Also sets the user in the context.
func CheckIfUserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := jwtutils.GetSubFromTokenFromContext(c)
		if err != nil {
			log.Println("Error getting sub from token: ", err.Error())
			c.AbortWithStatusJSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		user := models.User{
			Sub: sub,
		}
		err = user.Find()
		if err != nil {
			log.Println("Error finding user: ", err.Error())
			c.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
