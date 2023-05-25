package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/controllers"
	"github.com/tang-projects/api_go/internal/utils"
)

func setupUser(r *gin.Engine) {
	ctrl := controllers.UserController{}

	r.POST("api/v1/users", ctrl.CreateUser)
	r.GET("api/v1/users/:id", utils.AuthMiddleware, ctrl.ReadUser)
	r.PUT("api/v1//users/:id", utils.AuthMiddleware, ctrl.UpdateUser)
	r.DELETE("api/v1//users/:id", utils.AuthMiddleware, ctrl.DeleteUser) // 谨慎提供给用户，一般来说不会开放给用户进行操作
	r.POST("api/v1/users/logout", utils.AuthMiddleware, ctrl.LogoutUser)
}
