package router

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/utils"
)

func Run(port string) {
	r := gin.Default()

	// 配置 CORS 策略
	// 使用 os.Getenv 函数来获取 CORS_ALLOW_ORIGINS 环境变量的值，并使用 strings.Split 函数将其分割为一个字符串切片。然后我们将这个字符串切片赋值给 AllowOrigins 字段。在本地开发环境中，可以将 CORS_ALLOW_ORIGINS 环境变量设置为 "http://localhost:3000"，以允许来自本地前端应用程序的跨域请求。在生产环境中，可以将 CORS_ALLOW_ORIGINS 环境变量设置为 "http://cra.tangzhenming.com"，以允许来自生产环境前端应用程序的跨域请求。
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	r.Use(cors.New(config))

	setupUserRoutes(r)
	setupPostRoutes(r)

	r.Run(fmt.Sprintf(":%s", port))
}

// 路由 AuthMiddleware 权限控制
// Use a whitelist to optimize the code
func AddAuthRoute(r *gin.Engine, method string, path string, handler gin.HandlerFunc, auth bool) {
	if auth {
		r.Handle(method, path, utils.AuthMiddleware, handler)
	} else {
		r.Handle(method, path, handler)
	}
}
