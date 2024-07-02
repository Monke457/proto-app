package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const(
	userBase string = "/users" 
)

func AddControllers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "Mellow Yellow")
	})

	e.GET(userBase, func(c echo.Context) error {
		id := c.QueryParam("id")
		if id != "" {
			return c.Render(http.StatusOK, "users/index", id)
		}
		return c.Render(http.StatusOK, "users/index", "Users")
	})

	e.GET(userBase + "/edit", func(c echo.Context) error {
		id := c.QueryParam("id")
		if id != "" {
			return c.Render(http.StatusOK, "users/edit", id)
		}
		return c.Redirect(http.StatusPermanentRedirect, userBase)
	})
}
