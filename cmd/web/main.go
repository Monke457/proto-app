package main

import (
	"app/internal/env"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Name struct {
	Name string
}

func main() {
	err := env.Init()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	name := Name { Name: "Mellow Yellow" }

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", name)
	})

	port, _ := env.Get("APP_PORT")
	if port == "" {
		port = "1337"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
