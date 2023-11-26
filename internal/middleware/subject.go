package middleware

import (
	"educahub/internal/models"
	"log"

	gin "github.com/gin-gonic/gin"
)

// Checks if subject exists in the database.
// The path requires a community_url and a subject_url.
// Also sets the subject in the context.
func CheckIfSubjectExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		subjectUrl := c.Param("subject_url")
		subject := models.Subject{
			URL: subjectUrl,
		}
		err := subject.FindByURL()
		if err != nil {
			log.Println("Error finding subject: ", err.Error())
			c.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		community := c.MustGet("community").(models.Community)

		if subject.CommunityID != community.ID {
			log.Println("Subject not found in community")
			c.AbortWithStatusJSON(404, gin.H{
				"message": "subject not found",
			})
			return
		}

		c.Set("subject", subject)
		c.Next()
	}
}
