package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

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
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	name := Name { Name: "World" }

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", name)
	})

	port := os.Getenv("PORT")
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
