package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

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

// 发送邮件
func SendEmail(email string) (string, error) {
	randCode := GenerateRandCode() // 生成随机验证码

	m := gomail.NewMessage()

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("tangzhenming1207@qq.com", "Ryan")},
		"To":      {email},
		"Subject": {"XXX验证码"},
	})
	m.SetBody("text/html", fmt.Sprintf(`您正在登录/注册 XXX <br/> 您的验证码是：%s <br /> 验证码有效期为 24 小时`, randCode))

	d, err := createDialer()
	if err != nil {
		return "", err
	}

	// 配置 TLS 连接
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err = d.DialAndSend(m); err != nil {
		return "", err
	}

	return randCode, nil
}
