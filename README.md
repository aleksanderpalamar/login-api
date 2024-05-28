# Login API in Golang

This project is a sample implementation of a login API in Golang, using [Echo](https://echo.labstack.com/) and [JWT](https://github.com/appleboy/gin-jwt), and [GORM](https://gorm.io/) as the database.

## Features

- User registration (sign up)
- User authentication (sign in)
- User logout (sign out)
- JWT-based authentication

## Project Structure
```sh
rbac-api/
├── main.go
├── db/
│ └── db.go
├── handlers/
│ └── auth.go
├── models/
│ └── user.go
└── utils/
└── jwt.go
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/aleksanderpalamar/login-api.git
cd login-api
```

2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the applications: `./main`

## Running the Application

To start the application, run:

```sh
go run main.go
```

The server will start on http://localhost:3000.

**API Endpoints**
- Sign Up
- URL: /signup
- Method: POST
- Request Body:

```json
{
  "username": "testuser",
  "password": "testpass",
  "role": "user"
}
```
- Response:

```json
{
  "message": "User created successfully"
}
```
**Sign Out**
- URL: /signout
- Method: POST
- Response:

```json
{
  "message": "signed out"
}
```

## Utility Functions

JWT Utility
The JWT utility provides functions for generating and validating JWT tokens.

GenerateJWT
Generates a JWT token for a given username.

```go
func GenerateToken(username, password string) (string, error)
```

ValidateJWT
Validates a given JWT token.

```go
func ValidateJWT(tokenString string) (*Claims, error)
```

## Dependencies
Echo - High performance, minimalist Go web framework
JWT - JWT (JSON Web Tokens) implementation in Golang

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Credits

- Developer by [Aleksander Palamar](https://github.com/aleksanderpalamar)