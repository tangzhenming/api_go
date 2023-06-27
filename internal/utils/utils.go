package utils

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 生成严格均匀分布的 6 位随机数字
func GenerateRandCode() string {
	const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 6 // 数字长度
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			panic(err)
		}
		result[i] = letterBytes[num.Int64()]
	}
	return string(result)
}

// interface{} 是 Go 语言中的空接口类型。它可以表示任何类型的值
// 可以使用类型断言或类型切换来获取 data 参数的实际类型和值
// if str, ok := data.(string); ok {
// 	// data is a string
// 	fmt.Println(str)
// } else if num, ok := data.(int); ok {
// 	// data is an int
// 	fmt.Println(num)
// } else {
// 	// data is some other type
// }
// 类型断言的语法是 x.(T)，其中 x 是一个接口类型的值，T 是一个类型。类型断言会检查 x 是否持有一个 T 类型的值，并返回这个值和一个布尔值

// 替换 c.JSON 调用
func RespondJSON(c *gin.Context, code int, data interface{}, message string) {
	httpStatus, exists := c.Get("httpStatus") // 从上下文中获取，一些特殊的接口比如需要鉴权的接口，在调用 RespondJSON 之前需要提前设置好 httpStatus
	if !exists {
		httpStatus = http.StatusOK
	}
	c.JSON(httpStatus.(int), gin.H{"code": code, "data": data, "message": message})
}
