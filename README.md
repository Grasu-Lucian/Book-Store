# Book Store

Welcome to Project Book Store!

Book Store RESTful API for managing books in a store with added user authentication.

### Key Features:
1.User Registration and Login: The /register and /login routes handle user registration and login respectively.

2.Book Management: The /book and /books routes handle creating, retrieving, updating, and deleting books. The /:id in the /book/:id routes is a placeholder for the ID of a book.

3.Authentication Middleware: The middleware.RequireAuth function is used as middleware for the routes that modify books. This function checks if the user is authenticated before the request reaches the controller..


### Technology Stack:  
Backend:Go, GORM, bcrypt, Gin  
Database:PostgreSQL,  
Authentication: JWT (JSON Web Tokens), OAuth  

### Installation:
1. Clone the repository:
```Bash
 git clone https://github.com/username/project.git
```
2. To install the GORM package in your Go environment, use the following command:
```Bash
go get -u gorm.io/gorm
```
3.To install support for PostgreSQL to the project, you need to install the PostgreSQL driver using the following command:
```Bash
go get -u gorm.io/driver/postgres
```
4.To install the Gin framework to the project, you can use the following command:
```Bash
$ go get -u github.com/gin-gonic/gin
```
5. To install the bcrypt package use the following command:
```bash
go get -u golang.org/x/crypto/bcrypt
```

6. To install the jwt package use the following command:
```bash
go get -u github.com/golang-jwt/jwt/v5
```
7. To install the godotenv package use the following command:
```bash
go get github.com/joho/godotenv
```
7. To install the compile daemon package use the following command:
```bash
go get github.com/githubnemo/CompileDaemon
```
### Usage
In order to start the project you must run 
```bash
go run main.go
```
 You have the following routes that you can use:
 
- Register a new user: `curl -X POST -H "Content-Type: application/json" -d '{"user":"Damien", "password":"password"}' http://localhost:3000/register`
- Login: `curl -X POST -H "Content-Type: application/json" -d '{"user":"Damien", "password":"password"}' http://localhost:3000/login`
- Add a new book (requires authentication): `curl -X POST -H "Content-Type: application/json" -d '{"title":"Example Book", "author":"Example Author", "publishedDate":"2000-01-01", "isbn":"123-456-789", "price":9.99}' http://localhost:3000/book`
- Get a book by ID: `curl http://localhost:3000/book/:id`
- Get all books: `curl http://localhost:3000/books`
- Update a book by ID (requires authentication): `curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Book", "author":"Updated Author", "publishedDate":"2001-01-01", "isbn":"987-654-321", "price":19.99}' http://localhost:3000/book/:id`
- Delete a book by ID (requires authentication): `curl -X DELETE http://localhost:3000/book/:id`

