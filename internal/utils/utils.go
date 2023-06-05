package utils

import (
	"crypto/rand"
	"math/big"
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
