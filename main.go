package main

import (
	"github.com/aleksanderpalamar/login-api/db"
	"github.com/aleksanderpalamar/login-api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()

	e := echo.New()

	e.POST("/signup", handlers.SignUp)
	e.POST("/signin", handlers.SignIn)
	e.POST("/signout", handlers.SignOut)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}
