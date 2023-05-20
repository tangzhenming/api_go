# API-GO

Learning by AI chatbot.

- 2023.05.05
  1. Use [Gin](https://github.com/gin-gonic/gin) to handle routes
  2. Use [Gorm](https://github.com/go-gorm/gorm) and [PostgreSQL](https://www.postgresql.org/) to handle database
- 2023.05.19
  1. Use [Cobra](https://github.com/spf13/cobra) to custom some commands for command line usage
- 2023.05.20 
  1. Remove Gorm
  2. Write queries then use [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) to generate Go code that presents type-safe interfaces to those queries
  3. Use [golang-migrate](https://github.com/golang-migrate/migrate) tool to parsing migrations