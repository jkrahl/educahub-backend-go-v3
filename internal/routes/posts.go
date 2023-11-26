package routes

import (
	"educahub/internal/middleware"
	"educahub/internal/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func SetupPostsRoutes(r *gin.Engine) {
	AuthRequired, err := middleware.GetAuthMiddleware()
	if err != nil {
		panic(err)
	}

	r.GET("/communities", GetAllCommunitiesHandler)

	v1 := r.Group("/:community_url", middleware.CheckIfCommunityExists(), AuthRequired, middleware.CheckIfUserExists())
	{
		v1.GET("/", GetCommunityHandler)
		subjects := v1.Group("/subjects")
		{
			subjects.GET("/:subject_url", middleware.CheckIfSubjectExists(), GetSubjectHandler)
			subjects.GET("/:subject_url/posts", middleware.CheckIfSubjectExists(), GetAllPostsFromSubjectHandler)
			subjects.GET("/", GetAllSubjectsFromCommunityHandler)
		}
		posts := v1.Group("/posts")
		{
			posts.GET("/:post_url", middleware.CheckIfPostExists(), GetPostHandler)
			posts.DELETE("/:post_url", middleware.CheckIfPostExists(), DeletePostHandler)
			posts.GET("/", GetAllPostsFromCommunityHandler)
			posts.POST("/", CreatePostHandler)
			// Comments
			comments := posts.Group("/:post_url/comments", middleware.CheckIfPostExists())
			{
				comments.DELETE("/:comment_uuid", DeleteCommentHandler)
				comments.GET("/", GetAllPostCommentsHandler)
				comments.POST("/", CreateCommentHandler)
			}
		}
	}
}

func GetAllCommunitiesHandler(c *gin.Context) {
	communities, err := models.GetAllCommunities()
	if err != nil {
		log.Println("Error getting all communities: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response []models.CommunityResponse

	for _, community := range communities {
		response = append(response, models.CommunityToCommunityResponse(&community))
	}

	c.JSON(200, response)
}

func GetCommunityHandler(c *gin.Context) {
	community := c.MustGet("community").(models.Community)
	c.JSON(200, models.CommunityToCommunityResponse(&community))
}

func GetAllPostsFromCommunityHandler(c *gin.Context) {
	community := c.MustGet("community").(models.Community)

	posts, err := community.GetAllPosts()
	if err != nil {
		log.Println("Error getting all posts: ", err.Error())
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(posts) == 0 {
		c.JSON(200, []models.PostResponse{})
		return
	}

	var response []models.PostResponse

	for _, post := range posts {
		response = append(response, models.PostToPostResponse(&post))
	}

	c.JSON(200, response)
}

func GetPostHandler(c *gin.Context) {
	post := c.MustGet("post").(models.Post)

	c.JSON(200, models.PostToPostResponse(&post))
}

func CreatePostHandler(c *gin.Context) {
	community := c.MustGet("community").(models.Community)

	body := struct {
		Type       int    `json:"type" binding:"required,gte=1,lte=2" form:"type"`
		Title      string `json:"title" binding:"required" form:"title"`
		Content    string `json:"content" binding:"required" form:"content"`
		SubjectURL string `json:"subject_url" binding:"required" form:"subject_url"`
	}{}
	err := c.BindJSON(&body)
	if err != nil {
		log.Println("Error binding body: ", err.Error())
		c.JSON(400, gin.H{
			"message": "invalid body",
		})
		return
	}

	user := c.MustGet("user").(models.User)

	subject := models.Subject{
		URL: body.SubjectURL,
	}

	err = subject.FindByURL()
	if err != nil {
		log.Println("Error finding subject: ", err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	randomUUID := uuid.New()

	post := models.Post{
		Type:        body.Type,
		Title:       body.Title,
		Content:     body.Content,
		UserID:      user.ID,
		User:        user,
		CommunityID: community.ID,
		Community:   community,
		URL:         slug.Make(body.Title) + "-" + strings.Split(randomUUID.String(), "-")[0],
		SubjectID:   subject.ID,
		Subject:     subject,
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
		"url":     post.URL,
	})
}

func DeletePostHandler(c *gin.Context) {
	post := c.MustGet("post").(models.Post)
	user := c.MustGet("user").(models.User)

	if post.UserID != user.ID {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	err := post.Delete()
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
