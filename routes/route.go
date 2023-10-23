package routes

import (
	"main/get"
	"main/post"
	"main/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user", user.HandleUser)
	e.GET("/bio",  get.HandleUser)
	e.POST("/bio", post.HandleUser)
}