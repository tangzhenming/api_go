package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/tang-projects/api_go/internal/models"
	"github.com/tang-projects/api_go/internal/utils"
	"gorm.io/gorm"
)

type UserController struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

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
		ctrl.RedisClient.Set(user.Email, randCode, time.Minute*5) // 验证码有效期为 5 分钟

		c.JSON(http.StatusOK, gin.H{"message": "Verification code sent."})
		return
	}

	// 提供了验证码
	// 如果验证验证码失败，流程中止
	storedCode, err := ctrl.RedisClient.Get(user.Email).Result()
	if err != nil || user.VerificationCode != storedCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code"})
		return
	}

	// 验证通过，创建/更新用户帐户
	result := ctrl.DB.Where("Email = ?", user.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果帐户不存在，创建一个新帐户
		result = ctrl.DB.Create(&user)
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	} else {
		// 如果帐户已经存在，更新用户信息
		result = ctrl.DB.Model(&models.User{}).Where("Email = ?", user.Email).Updates(user)
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (ctrl UserController) ReadUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := ctrl.DB.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	result := ctrl.DB.Model(&models.User{}).Where("ID = ?", id).Updates(user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or no fields updated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (ctrl UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := ctrl.DB.Delete(&models.User{}, id)
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
