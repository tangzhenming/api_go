package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tang-projects/api_go/docs"
	"gorm.io/gorm"
)

//	@title			API GO
//	@version		1.0
//	@description	API GO 接口文档

//	@contact.name	Ryan
//	@contact.email	tangzhenming1207@gmail.com

//	@host	localhost:8080
func Run(DB *gorm.DB) {
	r := gin.Default()

	setupUser(r, DB)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
