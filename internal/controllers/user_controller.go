package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// CreateUser godoc
//	@Summary	Create a user
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Failure	500
//	@Router		/api/v1/users [post]
func (ctrl UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	result := ctrl.DB.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

// ReadUser godoc
//	@Summary	Read a user
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Failure	500
//	@Param		id	path	int	true	"UserID"
//	@Router		/api/v1/users/{id} [get]
func (ctrl UserController) ReadUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := ctrl.DB.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
//	@Summary	Update a user
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Failure	500
//	@Param		id	path	int	true	"UserID"
//	@Router		/api/v1/users/{id} [put]
func (ctrl UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	result := ctrl.DB.Model(&models.User{}).Where("user_id = ?", id).Updates(user)
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

// UpdateUser godoc
//	@Summary	Update a user
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Failure	500
//	@Param		id	path	int	true	"UserID"
//	@Router		/api/v1/users/{id} [delete]
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
