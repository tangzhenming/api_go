package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// 生成严格均匀分布的 6 位随机数字
func GenerateRandCode() string {
	const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 6 // 验证码长度
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

// 发送邮件
func SendEmail(email, randCode string) error {
	m := gomail.NewMessage()

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("tangzhenming1207@qq.com", "Ryan")},
		"To":      {email, "tangzhenming1207@gmail.com"},
		"Subject": {"XXX验证码"},
	})
	m.SetBody("text/html", fmt.Sprintf(`您正在登录/注册 XXX <br/> 您的验证码是：%s <br /> 验证码有效期为 5 分钟`, randCode))

	d, err := createDialer()
	if err != nil {
		return err
	}

	// Send the email
	if err = d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
func createDialer() (*gomail.Dialer, error) {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return gomail.NewDialer(host, port, username, password), nil
}
