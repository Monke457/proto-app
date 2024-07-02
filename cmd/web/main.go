package main

import (
	"app/internal/env"
	"app/internal/renderer"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	port := env.Get("APP_PORT")
	if port == "" {
		port = "1337"
	}

	e.Renderer = renderer.New()
	e.GET("/tickets", func(c echo.Context) error {
		return c.Render(http.StatusOK, "tickets/index", "Tickets")
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "Mellow Yellow")
	})

	e.Logger.Fatal(e.Start(":" + port))
}

