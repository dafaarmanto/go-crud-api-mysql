# Go CRUD API + MySQL Bookstore

Bookstore App CRUD API with Go Language + MySQL using jinzhu/gorm, go-sql-driver/mysql, gorilla/mux and Postman

#### Required: Go, MySQL, and Postman (or curl, insomnia, etc)

----
1. Open Terminal and go to project directory
2. Run MySQL, don't forget to create the database as you wish
3. Type `go mod init example.com/go-bookstore`
4. Install gorilla/mux, type `go get "github.com/gorilla/mux"`
5. Install jinzhu/gorm, type `go get "github.com/jinzhu/gorm"`
6. Install go-sql-driver/mysql, type `go get "github.com/jinzhu/gorm/dialects/mysql"`
7. Install godotenv, type `go get github.com/joho/godotenv`
8. Create .env file, please refer to .env.example file
9. Type `go run main.go` in terminal
10. Open Postman, and try to send some request

---
#### URI
> localhost:9010/book/

#### Requests
| Method | Endpoint   | Description        |
|--------|------------|--------------------|
| GET    | /book/     | Get All Books      |
| GET    | /book/{id} | Get Books by ID    |
| POST   | /book/     | Create             |
| PUT    | /book/{id} | Update             |
| DELETE | /book/{id} | Delete Books by ID |

```azure
{
    {
        "Name": "",
        "Author": "",
        "Publication": ""
    }
}
```
