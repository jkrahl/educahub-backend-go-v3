package routes

import (
	"educahub/internal/jwt"
	"educahub/internal/middleware"
	"educahub/internal/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func SetupPostsRoutes(r *gin.Engine) {
	jwtMiddleware, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/posts")
	{
		v1.GET("/", GetAllPostsHandler)
		v1.GET("/:uuid", GetPostHandler)
		v1.POST("/", jwtMiddleware, CreatePostHandler)
		v1.DELETE("/:uuid", jwtMiddleware, DeletePostHandler)
		// Comments
		v1.GET("/:uuid/comments", GetAllCommentsHandler)
		v1.POST("/:uuid/comments", jwtMiddleware, CreateCommentHandler)
		v1.DELETE("/:uuid/comments/:comment_uuid", jwtMiddleware, DeleteCommentHandler)
	}
}

func GetAllPostsHandler(c *gin.Context) {
	posts := []models.Post{}
	response := []models.PostResponse{}
	err := models.GetAllPosts(&posts)
	if err != nil {
		log.Println("Error getting all posts: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Convert posts to post responses
	for _, post := range posts {
		response = append(response, *models.PostToPostResponse(&post))
	}

	c.JSON(200, response)
}

func GetPostHandler(c *gin.Context) {
	postUuid := c.Param("uuid")
	post := models.Post{
		URL: postUuid,
	}
	err := post.Find()
	if err != nil {
		log.Println("Error finding post: ", err.Error())
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, models.PostToPostResponse(&post))
}

func CreatePostHandler(c *gin.Context) {
	sub, err := jwt.GetSubFromTokenFromContext(c)
	if err != nil {
		log.Println("Error getting sub from token: ", err.Error())
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}
	body := struct {
		Type    int    `json:"type" binding:"required,gte=1,lte=2" form:"type"`
		Title   string `json:"title" binding:"required" form:"title"`
		Content string `json:"content" binding:"required" form:"content"`
		Subject string `json:"subject" form:"subject"`
		Unit    string `json:"unit" form:"unit"`
	}{}
	err = c.BindJSON(&body)
	if err != nil {
		log.Println("Error binding body: ", err.Error())
		c.JSON(400, gin.H{
			"message": "invalid body",
		})
		return
	}

	user := models.User{
		Sub: sub,
	}
	err = user.Find()
	if err != nil {
		log.Println("Error finding user: ", err.Error())
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	if body.Type == models.PostTypeNotes {
		// TODO: Send to S3
		body.Content = "NOTES"
	}

	randomUUID := uuid.New()

	post := models.Post{
		Type:    body.Type,
		Title:   body.Title,
		Content: body.Content,
		UserID:  user.ID,
		User:    user,
		URL:     slug.Make(body.Title) + "-" + strings.Split(randomUUID.String(), "-")[0],
		Subject: body.Subject,
		Unit:    body.Unit,
	}

	err = post.Create()
	if err != nil {
		log.Println("Error creating post: ", err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "post created",
		"uuid":    post.URL,
	})
}

func DeletePostHandler(c *gin.Context) {
	sub, err := jwt.GetSubFromTokenFromContext(c)
	if err != nil {
		log.Println("Error getting sub from token: ", err.Error())
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	postUuid := c.Param("uuid")
	post := models.Post{
		URL: postUuid,
	}
	err = post.Find()
	if err != nil {
		log.Println("Error finding post: ", err.Error())
		c.JSON(404, gin.H{
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
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	if post.UserID != user.ID {
		log.Println("Error deleting post: ", err.Error())
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	err = post.Delete()
	if err != nil {
		log.Println("Error deleting post: ", err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "post deleted",
	})
}
