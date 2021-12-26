package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucas-simao/go-basic-api/user"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	user.RouterGroup(e)

	e.Logger.Fatal(e.Start(":9000"))
}
