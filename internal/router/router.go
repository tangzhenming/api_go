package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	r := gin.Default()

	setupUser(r)

	r.Run(fmt.Sprintf(":%s", port))
}
