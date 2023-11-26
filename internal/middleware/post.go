package middleware

import (
	"educahub/internal/models"
	"log"

	gin "github.com/gin-gonic/gin"
)

// This function is used to check if a post exists in the database in a community.
// The path requires a community_url and a post_url.
// Also sets the post in the context.
func CheckIfPostExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		postUrl := c.Param("post_url")
		post := models.Post{
			URL: postUrl,
		}
		err := post.FindByURL()
		if err != nil {
			log.Println("Error finding post: ", err.Error())
			c.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		community := c.MustGet("community").(models.Community)

		if post.CommunityID != community.ID {
			log.Println("Post not found in community")
			c.AbortWithStatusJSON(404, gin.H{
				"message": "post not found",
			})
			return
		}

		log.Println("Key set: Post")
		c.Set("post", post)
		c.Next()
	}
}
