package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
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
	m.SetBody("text/html", fmt.Sprintf(`您正在登录/注册 XXX <br/> 您的验证码是：%s <br /> 验证码有效期为 30 分钟`, randCode))

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

// 生成 JWT
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

var jwtSecret = []byte("my_secret_key") // my_secret_key 是一个固定的字符串密钥，建议使用 ASCII 字符作为密钥，因为它们更容易在不同的系统和编程语言之间传输和处理，比如 my_secret_key_123

// 接受一个 userID 参数，并返回一个包含该用户 ID 的 JWT
func GenerateToken(userID uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my_project", // 表示 JWT 的签发者，可以使用任何字符串，包括中文
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "请先登录"})
		c.Abort() // gin 框架中 Context 类型的一个方法，它用于终止当前请求的处理。当你在中间件或处理函数中调用 c.Abort() 时，后续的中间件和处理函数都不会被执行，请求的处理将立即结束
		return
	}

	claims, err := parseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Token 无效"})
		c.Abort()
		return
	}

	// 查询数据库，判断 Token 是否已经被标记为空（用户注销登录）
	var user models.User
	if result := db.PG.Where("id = ? AND token = ?", claims.UserID, token).First(&user); result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Token 已失效"})
		c.Abort()
		return
	}

	// gin 框架中 Context 类型的一个方法，它用于在上下文中存储键值对。你可以在处理请求的过程中使用 c.Set 来存储一些数据，这些数据可以在后续的中间件和处理函数中通过 c.Get 方法来获取；例如 c.Set("userID", claims.UserID) 将 userID 存储在 JWT 中，后续如果接口只允许用户访问自己的信息，那么可以不使用 URL 参数来传递，而是直接使用 Get 方法获取 userID 后使用，但如果接口允许其他人/管理员访问当前用户信息，那就依然要使用 URL 参数来处理 userID ，比如社区帖子中点击别人的头像查看别人的信息
	c.Set("userID", claims.UserID)
}
func parseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}
