package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
	"github.com/tang-projects/api_go/internal/utils"
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
		utils.RespondJSON(c, 0, nil, err.Error())
		return
	}

	// 检查用户是否存在
	var user models.User
	result := db.PG.First(&user, input.UserID)
	if result.Error == gorm.ErrRecordNotFound {
		utils.RespondJSON(c, 0, nil, "User not found")
		return
	} else if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}
	result = db.PG.Create(&post)
	if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}

	// 查询关联的用户对象并将其存储到 post.User 字段中
	err := db.PG.Model(&post).Association("User").Find(&post.User)
	if err != nil {
		utils.RespondJSON(c, 0, nil, err.Error())
		return
	}

	utils.RespondJSON(c, 1, post, "")
}

func (ctrl PostController) ReadPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, 0, nil, "Invalid post ID")
		return
	}

	var post models.Post
	result := db.PG.Preload("User").First(&post, id)
	if result.Error == gorm.ErrRecordNotFound {
		utils.RespondJSON(c, 0, nil, "Post not found")
		return
	} else if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}

	utils.RespondJSON(c, 1, post, "")
}

func (ctrl PostController) ReadPosts(c *gin.Context) {
	var input struct {
		StartTime string `form:"start_time"`
		EndTime   string `form:"end_time"`
		Page      int    `form:"page"`
		PageSize  int    `form:"page_size"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		utils.RespondJSON(c, 0, nil, err.Error())
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

	var total int64
	result := query.Count(&total)
	if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}
	// 使用 Gorm 的 Offset 和 Limit 方法来实现分页查询
	result = query.Offset((input.Page - 1) * input.PageSize).Limit(input.PageSize).Find(&posts)
	if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}

	responsePosts := make([]struct {
		models.Post
		User models.ResponseUser `json:"User"`
	}, len(posts))
	for i, post := range posts {
		responsePosts[i].Post = post
		responsePosts[i].User = models.ResponseUser{
			ID:        post.User.ID,
			CreatedAt: post.User.CreatedAt,
			UpdatedAt: post.User.UpdatedAt,
			Name:      post.User.Name,
			Email:     post.User.Email,
		}
	}

	data := map[string]interface{}{
		"data":  responsePosts,
		"total": total,
	}
	utils.RespondJSON(c, 1, data, "")
}

func (ctrl PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondJSON(c, 0, nil, "Invalid post ID")
		return
	}

	result := db.PG.Delete(&models.Post{}, id)
	if result.Error == gorm.ErrRecordNotFound {
		utils.RespondJSON(c, 0, nil, "Post not found")
		return
	} else if result.Error != nil {
		utils.RespondJSON(c, 0, nil, result.Error.Error())
		return
	}

	utils.RespondJSON(c, 1, nil, "deleted")
}
