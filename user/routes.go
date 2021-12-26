package user

import (
	"github.com/labstack/echo/v4"
)

func RouterGroup(e *echo.Echo) {

	g := e.Group("/users")

	g.GET("/", getUserHandler)
	g.POST("/new", postUserHandler)
}
