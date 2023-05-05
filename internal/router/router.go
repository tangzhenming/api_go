package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/controller"
)

func RunServe() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
