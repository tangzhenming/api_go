package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/utils"
)

func Run(port string) {
	r := gin.Default()

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
