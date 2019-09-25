# Go Lang Starter Pack

This is golang starter pack using MVC pattern

## About

This is golang starter pack using MVC pattern

- Focus on REST API Building
- Support HTTP Verb (GET, POST, PUT, PATCH, DELETE)
- Support multiple middlewares and specific route middlewares
- JWT (JSON Web Token) out of the box

#### Download Dependecies

`$ go get github.com/joho/godotenv`

`$ go get github.com/gorilla/mux`

`$ go get github.com/gorilla/context`

`$ go get github.com/go-sql-driver/mysql`

`$ go get github.com/dgrijalva/jwt-go`

`$ go get golang.org/x/crypto/bcrypt`

#### Setup

`$ copy .env.example to .env`

`$ setup database config inside .env file (databse, username, password)`

`> Import database from databases folder`

#### Run Apps

`$ go run main.go`
 