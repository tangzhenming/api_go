package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
	"gorm.io/gorm"
)

type PostController struct{}

func (ctrl PostController) CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 检查用户是否存在
	var user models.User
	result := db.PG.First(&user, input.UserID)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}
	result = db.PG.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 查询关联的用户对象并将其存储到 post.User 字段中
	err := db.PG.Model(&post).Association("User").Find(&post.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

func (ctrl PostController) ReadPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	result := db.PG.Preload("User").First(&post, id)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

type ResponseUser struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
}

func (ctrl PostController) ReadPostsByTimeRange(c *gin.Context) {
	var input struct {
		StartTime string `form:"start_time"`
		EndTime   string `form:"end_time"`
		Page      int    `form:"page"`
		PageSize  int    `form:"page_size"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters", "details": err.Error()})
		return
	}

	if input.Page <= 0 {
		input.Page = 1
	}
	if input.PageSize <= 0 {
		input.PageSize = 10
	}

	var posts []models.Post
	query := db.PG.Preload("User", func(db *gorm.DB) *gorm.DB { // 使用 Gorm 的 Preload 方法来自动加载关联的用户对象
		return db.Select("id", "created_at", "updated_at", "name", "email") // 使用 Gorm 的 Select 方法来指定返回的字段，从而避免返回敏感信息
	}).Model(&models.Post{})
	if input.StartTime != "" && input.EndTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}
	// 使用 Gorm 的 Offset 和 Limit 方法来实现分页查询
	result := query.Offset((input.Page - 1) * input.PageSize).Limit(input.PageSize).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	responsePosts := make([]struct {
		models.Post
		User ResponseUser `json:"User"`
	}, len(posts))
	for i, post := range posts {
		responsePosts[i].Post = post
		responsePosts[i].User = ResponseUser{
			ID:        post.User.ID,
			CreatedAt: post.User.CreatedAt,
			UpdatedAt: post.User.UpdatedAt,
			Name:      post.User.Name,
			Email:     post.User.Email,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": responsePosts})
}

func (ctrl PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	result := db.PG.Delete(&models.Post{}, id)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
