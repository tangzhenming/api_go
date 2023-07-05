package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
	"github.com/tang-projects/api_go/internal/utils"
	"gorm.io/gorm"
)

type UserController struct{}

// CreateUser 方法用于创建或登录用户账户。它接收一个 JSON 格式的请求体，其中包含用户的电子邮件地址和验证码。如果请求体中没有提供验证码，那么它会通过邮件发送验证码给用户。如果请求体中提供了验证码，那么它会验证验证码是否正确，然后根据电子邮件地址创建或登录用户账户。
func (ctrl UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, 0, nil, "Invalid request body")
		return
	}

	// 如果没有提供验证码，则通过邮件发送验证码，流程中止
	if user.VerificationCode == "" {

		randCode, err := utils.SendEmail(user.Email)
		if err != nil {
			utils.RespondJSON(c, 0, nil, err.Error())
			return
		}

		// 缓存验证码和电子邮件地址
		db.RedisClient.Set(user.Email, randCode, time.Minute*60*24) // 验证码（ Redis 缓存）有效期为 24 小时

		utils.RespondJSON(c, 1, nil, "Verification code sent.")
		return
	}

	// 提供了验证码
	// 如果验证验证码失败，流程中止
	storedCode, err := db.RedisClient.Get(user.Email).Result()
	if err != nil || user.VerificationCode != storedCode {
		utils.RespondJSON(c, 0, nil, "Invalid verification code")
		return
	}

	// 验证通过，创建或登录用户帐户
	var (
		message string
	)

	// Find user with email
	result := db.PG.Unscoped().Where("Email = ?", user.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果帐户不存在，创建一个新帐户
		result = db.PG.Create(&user)
		if result.Error != nil {
			log.Println(result.Error)
			utils.RespondJSON(c, 0, nil, "Failed to create user")
			return
		}

		message = "Created"
	} else if result.Error != nil {
		log.Println(result.Error)
		utils.RespondJSON(c, 0, nil, "Failed to find user")
		return
	}

	// User found
	message = "Logined"

	// Reset token and deleted_at field
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.RespondJSON(c, 0, nil, "Failed to generate Token")
		return
	}
	result = db.PG.Unscoped().Model(&user).Updates(map[string]interface{}{"Token": token, "deleted_at": nil})
	if result.Error != nil {
		log.Println(result.Error)
		utils.RespondJSON(c, 0, nil, "Failed to reset token and deleted_at field in user")
		return
	}

	// 自定义响应数据结构
	responseUser := map[string]interface{}{
		"ID":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
		"Token": user.Token,
	}

	utils.RespondJSON(c, 1, responseUser, message)
}

// ReadUser 方法用于读取用户信息。它接收一个 URL 参数 id，表示要查询的用户ID。它会根据这个 ID 查询对应的用户信息，并返回查询结果。
func (ctrl UserController) ReadUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	result := db.PG.Omit("Token").First(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		utils.RespondJSON(c, 0, nil, "Failed to read user")
		return
	}
	if result.RowsAffected == 0 {
		utils.RespondJSON(c, 0, nil, "User not found or no fields readed")
		return
	}

	utils.RespondJSON(c, 1, user, "Readed")
}

// UpdateUser 方法用于更新用户信息。它接收一个 URL 参数 id，表示要更新的用户ID。它还接收一个 JSON 格式的请求体，其中包含要更新的字段及其新值。它会根据这些数据更新对应的用户信息。
func (ctrl UserController) UpdateUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	userID, _ := c.Get("userID")
	if id != fmt.Sprint(userID) {
		utils.RespondJSON(c, 0, nil, "Failed to update other user")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, 0, nil, err.Error())
		return
	}

	result := db.PG.Model(&models.User{}).Where("ID = ?", id).Updates(user)
	if result.Error != nil {
		log.Println(result.Error)
		utils.RespondJSON(c, 0, nil, "Failed to update user")
		return
	}
	if result.RowsAffected == 0 {
		utils.RespondJSON(c, 0, nil, "User not found or no fields updated")
		return
	}

	utils.RespondJSON(c, 1, nil, "Updated")
}

// DeleteUser 方法用于删除用户账户。它接收一个 URL 参数 id，表示要删除的用户ID。它会根据这个 ID 删除对应的用户账户。
func (ctrl UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	if id != fmt.Sprint(userID) {
		utils.RespondJSON(c, 0, nil, "Failed to delete other user")
		return
	}

	result := db.PG.Delete(&models.User{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		utils.RespondJSON(c, 0, nil, "Failed to delete user")
		return
	}
	if result.RowsAffected == 0 {
		utils.RespondJSON(c, 0, nil, "User not found")
		return
	}

	utils.RespondJSON(c, 1, nil, "deleted")
}

// LogoutUser 方法用于注销当前登录的用户。它不需要任何参数，只需调用这个方法即可将当前登录的用户注销。
func (ctrl UserController) LogoutUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	// 将 Token 标记为无效
	db.PG.Model(&models.User{}).Where("id = ?", userID).Update("Token", "")

	utils.RespondJSON(c, 1, nil, "Logout successfull")
}
