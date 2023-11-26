package routes

import (
	"educahub/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetSubjectHandler(c *gin.Context) {
	subject := c.MustGet("subject").(models.Subject)
	c.JSON(200, models.SubjectToSubjectResponse(&subject))
}

func GetAllPostsFromSubjectHandler(c *gin.Context) {
	subject := c.MustGet("subject").(models.Subject)

	posts, err := subject.GetAllPosts()
	if err != nil {
		log.Println("Error getting all posts from subject: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(posts) == 0 {
		c.JSON(200, []models.PostResponse{})
		return
	}

	response := []models.PostResponse{}
	for _, post := range posts {
		response = append(response, models.PostToPostResponse(&post))
	}

	c.JSON(200, response)
}

// Community middleware already preloads the subjects.
func GetAllSubjectsFromCommunityHandler(c *gin.Context) {
	community := c.MustGet("community").(models.Community)

	subjects := community.Subjects

	if len(subjects) == 0 {
		c.JSON(200, []models.SubjectResponse{})
		return
	}

	response := []models.SubjectResponse{}
	for _, subject := range subjects {
		response = append(response, models.SubjectToSubjectResponse(&subject))
	}

	c.JSON(200, response)
}
