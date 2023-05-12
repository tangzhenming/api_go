package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(DB *gorm.DB) {
	r := gin.Default()

	setupUser(r, DB)

	r.Run()
}
