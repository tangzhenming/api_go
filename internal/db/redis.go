package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedis() {
	var client *redis.Client

	// 创建一个 Redis 客户端
	client = redis.NewClient(&redis.Options{
		Addr:     "api-go-redis:6379",
		Password: "",
		DB:       0,
	})

	// 测试连接是否正常
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	RedisClient = client

	// // 设置一个键值对
	// err = client.Set("mykey", "myvalue", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// // 获取一个键的值
	// val, err := client.Get("mykey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("mykey", val)

	// // 获取一个不存在的键的值
	// val2, err := client.Get("mykey2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("mykey2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("mykey2", val2)
	// }
}
