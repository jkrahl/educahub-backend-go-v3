package routes

import (
	"educahub/internal/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPostCommentsHandler(c *gin.Context) {
	post := c.MustGet("post").(models.Post)

	comments, err := post.GetAllComments()
	if err != nil {
		log.Println("Error getting all comments: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(comments) == 0 {
		c.JSON(200, []models.CommentResponse{})
		return
	}

	response := []models.CommentResponse{}
	for _, comment := range comments {
		response = append(response, *models.CommentToCommentResponse(&comment))
	}

	c.JSON(200, response)
}

func CreateCommentHandler(c *gin.Context) {
	body := struct {
		Content string `json:"content" binding:"required"`
	}{}
	err := c.BindJSON(&body)
	if err != nil {
		log.Println("Error binding json: ", err.Error())
		c.JSON(500, gin.H{
			"message": "invalid body",
		})
		return
	}

	post := c.MustGet("post").(models.Post)

	user := c.MustGet("user").(models.User)

	randomUUID := uuid.New().String()

	comment := models.Comment{
		UUID:    randomUUID,
		Content: body.Content,
		UserID:  user.ID,
		User:    user,
		PostID:  post.ID,
		Post:    post,
	}

	err = comment.Create()
	if err != nil {
		log.Println("Error creating comment: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "comment created",
		"uuid":    comment.UUID,
	})
}

func DeleteCommentHandler(c *gin.Context) {
	post := c.MustGet("post").(models.Post)
	user := c.MustGet("user").(models.User)

	commentUuid := c.Param("comment_uuid")
	comment := models.Comment{
		UUID: commentUuid,
	}
	err := comment.Find()
	if err != nil {
		log.Println("Error finding comment: ", err.Error())
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	if comment.UserID != user.ID {
		log.Println("Error user not owner of comment")
		c.JSON(401, gin.H{
			"message": "user not owner of comment",
		})
		return
	}

	if comment.PostID != post.ID {
		log.Println("Error post not owner of comment")
		c.JSON(401, gin.H{
			"message": "comment not in post",
		})
		return
	}

	err = comment.Delete()
	if err != nil {
		log.Println("Error deleting comment: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "comment deleted",
	})
}
