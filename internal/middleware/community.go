package middleware

import (
	"educahub/internal/models"
	"log"

	gin "github.com/gin-gonic/gin"
)

// This function is used to check if a community exists in the database.
// The path requires a community_url.
// Also sets the community in the context.
func CheckIfCommunityExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		communityUrl := c.Param("community_url")
		community := models.Community{
			URL: communityUrl,
		}
		err := community.Find()
		if err != nil {
			log.Println("Error finding community: ", err.Error())
			c.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Set("community", community)
		c.Next()
	}
}
