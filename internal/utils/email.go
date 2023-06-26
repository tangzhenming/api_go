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
	// TLS（传输层安全协议）是一种用于在计算机网络上提供安全通信的协议。它可以在客户端和服务器之间建立一个安全的加密连接，以保护传输的数据不被窃听或篡改。
	// 当您的应用程序需要与其他服务器（例如电子邮件服务器、Web 服务器或 API 服务器）建立安全连接时，您需要配置 TLS 连接。配置 TLS 连接可以让您的应用程序验证服务器证书的有效性，并使用加密算法来保护传输的数据。

	// Send the email
	if err = d.DialAndSend(m); err != nil {
		return "", err
	}

	return randCode, nil
}
