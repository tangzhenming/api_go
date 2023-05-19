package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(DB *gorm.DB, port string) {
	r := gin.Default()

	setupUser(r, DB)

	r.Run(fmt.Sprintf(":%s", port))
}
