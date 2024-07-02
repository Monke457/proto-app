package renderer 

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	template *template.Template
}

func (h *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return h.template.ExecuteTemplate(w, name, data)
}

func New() *Renderer {
	t:= template.Must(template.ParseGlob("views/*.html"))
	template.Must(t.ParseGlob("views/**/*.html"))

	return &Renderer{ template: t }
}
