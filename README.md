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
