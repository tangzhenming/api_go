package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/controllers"
)

func setupUserRoutes(r *gin.Engine) {
	ctrl := controllers.UserController{}

	// One way to reduce the repetition of the r parameter and the common path is to defined a closure within this function that captures the parameters and paths and calls the AddAuthRoute function with it
	addAuthUserRoute := func(method string, path string, handler gin.HandlerFunc, auth bool) {
		basePath := "api/v1/users"
		AddAuthRoute(r, method, basePath+path, handler, auth)
	}
	addAuthUserRoute("POST", "", ctrl.CreateUser, false)
	addAuthUserRoute("GET", "/:id", ctrl.ReadUser, true)
	addAuthUserRoute("PUT", "/:id", ctrl.UpdateUser, true)
	addAuthUserRoute("DELETE", "/:id", ctrl.DeleteUser, true) // 谨慎提供给用户，一般来说不会开放给用户进行操作
	addAuthUserRoute("POST", "/logout", ctrl.LogoutUser, true)
}
