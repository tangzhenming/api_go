# API-GO

Learning by AI chatbot.

- **2023.05.05**
  - Use [Gin](https://github.com/gin-gonic/gin) to handle routes
  - Use [Gorm](https://github.com/go-gorm/gorm) and [PostgreSQL](https://www.postgresql.org/) to handle database
- **2023.05.12**
  - Use [joho/godotenv](https://github.com/joho/godotenv) to manage sensitive data
- **2023.05.19**
  - Use [Cobra](https://github.com/spf13/cobra) to custom some commands for command line usage
- **2023.05.23**
  - Use [crypto/rand](https://pkg.go.dev/crypto/rand) to generate a random verification code
  - Use [gomail](https://pkg.go.dev/gopkg.in/gomail.v2?utm_source=godoc#example-package) to send verification code to user
  - Use [Redis](https://redis.io/docs/getting-started/) to cache verification code: `docker run --name api-go-redis --network api-go-network -d redis`
- **2023.05.25**
  - 使用 JWT 验证用户身份的基本步骤
    1. 在用户登录后，服务端生成一个 JWT 并发送给客户端
    2. 客户端在之后的请求中都会带上这个 JWT ，服务端通过验证这个 JWT 来确认用户身份
  - 更多功能实现与细节
    1. 在需要验证用户身份的接口中，使用中间件来验证客户端传来的 token
    2. 创建或登录用户账户时都会新建 token ，刷新 token 时效性
    3. 提供注销登录接口，标记 token 为空
    4. 注销登录接口只能由当前用户操作，所以需要在鉴权中间件时将 userID 存储到路由上下文中，在 LogoutUser 控制器中通过路由上下文获取 userID 后标记其对应的 token 为空
- **2023.06.07**
  - 新增 Post 业务模块
    1. 声明 Post 模型
    2. 添加 migrate 和 drop 命令；执行命令使用 AutoMigrate 创建数据表
    3. 确认数据表创建完成后，即可使用 Post 模型去操作数据库
