package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/controllers"
)

func setupUser(r *gin.Engine) {
	ctrl := controllers.UserController{}

	r.POST("api/v1/users", ctrl.CreateUser)
	r.GET("api/v1/users/:id", ctrl.ReadUser)
	r.PUT("api/v1//users/:id", ctrl.UpdateUser)
	r.DELETE("api/v1//users/:id", ctrl.DeleteUser)
}
