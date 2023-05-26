package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
	"github.com/tang-projects/api_go/internal/utils"
	"gorm.io/gorm"
)

type UserController struct{}

// 通过邮箱验证码创建或登录用户账户
func (ctrl UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 如果没有提供验证码，则生成并发送验证码，流程中止
	if user.VerificationCode == "" {
		randCode := utils.GenerateRandCode() // 生成随机验证码

		err := utils.SendEmail(user.Email, randCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 缓存验证码和电子邮件地址
		db.RedisClient.Set(user.Email, randCode, time.Minute*30) // 验证码（ Redis 缓存）有效期为 30 分钟

		c.JSON(http.StatusOK, gin.H{"message": "Verification code sent."})
		return
	}

	// 提供了验证码
	// 如果验证验证码失败，流程中止
	storedCode, err := db.RedisClient.Get(user.Email).Result()
	if err != nil || user.VerificationCode != storedCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// 验证通过，创建或登录用户帐户
	var (
		code    int
		message string
	)
	result := db.PG.Where("Email = ?", user.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果帐户不存在，创建一个新帐户
		result = db.PG.Create(&user)
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		code = http.StatusCreated
		message = "Created"
	} else if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	} else {
		// 登录账户
		code = http.StatusOK
		message = "Logined"
	}

	// 无论是创建账户成功，还是登录账户成功，都重新生成 Token 并更新数据库
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Token"})
		return
	}
	db.PG.Model(&user).Update("Token", token)

	c.JSON(code, gin.H{"data": user, "message": message})
}

func (ctrl UserController) ReadUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	result := db.PG.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or no fields readed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Readed", "data": user})
}

func (ctrl UserController) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	result := db.PG.Model(&models.User{}).Where("ID = ?", id).Updates(user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or no fields updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (ctrl UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	if id != fmt.Sprint(userID) {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Failed to delete other user"})
		return
	}

	result := db.PG.Delete(&models.User{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (ctrl UserController) LogoutUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	// 将 Token 标记为无效
	db.PG.Model(&models.User{}).Where("id = ?", userID).Update("Token", "")

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Logout successfull"})
}
