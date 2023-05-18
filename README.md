# api_go

[Go Dev Doc](https://go.dev/doc/)

[Writing a Fully Fledged REST API Using PostgreSQL, Gin, and Gorm](https://betterprogramming.pub/writing-a-fully-fledged-api-using-postgresql-gin-and-gorm-4d5ba73114da)

[Build a REST API with Golang from scratch: PostgreSQL with Gorm and Gin Web Framework](https://dc1888.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7)

[Building a REST API with Golang using Gin and Gorm](https://blog.logrocket.com/rest-api-golang-gin-gorm/)

[李文周的博客](https://www.liwenzhou.com/posts/Go/golang-menu/)

## Go mod

`go mod init <project name>`

`go mod init github.com/tang-projects/api_go`

`go mod tidy`

[golang 系列: go mod 使用](https://zhuanlan.zhihu.com/p/394103953)

## Gin

[Go Gin](https://gin-gonic.com/docs/)

## PostgreSQL

- 本地新启一个 Docker 容器来运行数据库 `docker run -d --name api-go-pg -e POSTGRES_USER=api-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=api-go-pg-dev -e PGDATA=/var/lib/postgresql/data/pgdata -v api-go-pg-data:/var/lib/postgresql/data --network=api-go-network postgres:14`
- `docker network create api-go-network`
- `docker network connect api-go-network b63e`
- `docker start api-go-pg`
- 进入 PG container `docker exec -it <container name/id> bash`；或者使用 `ckolkman.vscode-postgres` 插件
- 连接数据库 `psql -U api-go -d api-go-dev`
- `\l` 查看所有数据库
- `\d` 查看数据库表 `\d <table name>` 查看单个数据库表详情
- `\q` 退出 PG 交互程序
- `exit` 退出 PG container

### 我是如何利用 AI 解决连接 PostgreSQL 过程中产生的问题的

1. 首先，我是基于 Docker 搭建了一个 Linux 系统环境，并在其中配置了常用的 Node/Ruby/Go 等工具；并搭配 Dev Containers 实现跨平台的开发环境
2. 直接启用新的 Docker 容器运行 PostgreSQL 数据库
3. 我询问 AI 在使用 Gin 的情况下如何连接 PostgreSQL ，并直接使用其给出的代码
   1. prompt: `I don't need a model, just give me a easiest way to show how to connect a postgresql database with Golang using the Gin framework`
   2. code:
   ```go
      package main

      import (
         "database/sql"
         "fmt"
         "log"

         "github.com/gin-gonic/gin"
         _ "github.com/lib/pq"
      )

      const (
         host     = "localhost"
         port     = 5432
         user     = "postgres"
         password = "your-password"
         dbname   = "your-dbname"
      )

      func main() {
         // Set up the database connection
         psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
            "password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)
         db, err := sql.Open("postgres", psqlInfo)
         if err != nil {
            log.Fatal(err)
         }
         defer db.Close()

         // Test the database connection
         err = db.Ping()
         if err != nil {
            log.Fatal(err)
         }
         fmt.Println("Successfully connected to the database!")

         // Set up the Gin router
         r := gin.Default()

         // Add a route that queries the database and returns the result
         r.GET("/ping", func(c *gin.Context) {
            var result string
            err := db.QueryRow("SELECT 'pong'").Scan(&result)
            if err != nil {
                  log.Fatal(err)
            }
            c.JSON(200, gin.H{
                  "message": result,
            })
         })

         // Start the Gin server
         r.Run()
      }
   ```
4. 直接运行，报错 `tcp 127.0.0.1:5432: connect: connection refused exit status 1` ，AI 给出提示需要检查 host, port, user, password, and dbname ，我确认以上信息无误并追问，AI 提示需要将两个 container 连接到同一个 network 才能实现通信
5. 我追问 AI ，如何确认这两个 container 在一个 network 中，AI 给出步骤
   1. Open a terminal and run the `docker network ls` command to list all the networks in your Docker environment.
   2. Find the name or ID of the network that you want to inspect. This should be the network that your containers are attached to.
   3. Run the `docker network inspect <network>` command, replacing `<network>` with the name or ID of the network you want to inspect.
   4. In the output of the `docker network inspect` command, look for a section named "Containers". This section lists all the containers that are attached to the network.
6. 发现 Dev Container 自动运行的 container 并未接入 PostgreSQL container 所接入的 network ，追问 AI 如何接入，AI 给出步骤
   1. Open a terminal and run the `docker ps` command to list all the running containers in your Docker environment.
   2. Find the name or ID of the PostgreSQL container and the name or ID of the Visual Studio Code dev container.
   3. Run the `docker inspect <container>` command, replacing `<container>` with the name or ID of the PostgreSQL container.
   4. In the output of the `docker inspect` command, look for a section named "NetworkSettings". This section contains information about the container’s network settings.
   5. In the "NetworkSettings" section, find the "Networks" field. This field lists all the networks that the container is attached to.
   6. Choose one of the networks from the "Networks" field and make a note of its name.
   7. Run the `docker network connect <network> <container>` command, replacing `<network>` with the name of the network you chose and `<container>` with the name or ID of the Visual Studio Code dev container.
7. 接入同一个 network 成功，依然报错 “dial tcp 127.0.0.1:5432: connect: connection refused exit status 1” ，追问 AI ，其再次让我检查 host
   1. answer: When running a PostgreSQL server in a Docker container, you cannot use localhost or 127.0.0.1 as the hostname to connect to the server from another container. Instead, you need to use the name of the PostgreSQL container as the hostname.
   2. 将 go 代码中的 host 设置为 PostgreSQL container name 后，`go run .` 成功运行

## MySQL

`docker run -d --name api-go-mysql -e MYSQL_DATABASE=api-go-mysql-dev -e MYSQL_USER=api-go -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v api-go-mysql-data:/var/lib/mysql --network=api-go-network mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci`

- `mysql -u api-go -p api-go-mysql-dev`
- `show databases;`
- `use api-go-mysql-dev;`
- `show tables;`
- `describe users;`

[Container 启动后又退出](https://stackoverflow.com/questions/40761876/cannot-restart-the-mysql-docker-container-gives-errors-like-cant-open-the-mys)

如果 container 运行异常，可使用 `docker logs <container name/id>` 来查看日志。

## 数据库工具

### [database/sql](https://pkg.go.dev/database/sql)

### [GORM](https://gorm.io/)

### [kyleconroy/sqlc]

## 命令行参数解析

[使用 os.Args 和 flag](https://www.liwenzhou.com/posts/Go/flag/)

[使用 cobra](https://github.com/spf13/cobra)