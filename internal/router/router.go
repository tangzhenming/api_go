package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func Run(RedisClient *redis.Client, DB *gorm.DB, port string) {
	r := gin.Default()

	setupUser(r, DB, RedisClient)

	r.Run(fmt.Sprintf(":%s", port))
}
