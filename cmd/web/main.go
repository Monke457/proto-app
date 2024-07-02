package main

import (
	"app/internal/db"
	"app/internal/env"
	"app/internal/controller"

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

	db.GenerateData()

	e.Renderer = controller.New()

	controller.AddControllers(e)

	e.Logger.Fatal(e.Start(":" + port))
}

