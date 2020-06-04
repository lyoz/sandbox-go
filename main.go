package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/lyoz/sandbox-go/handler"
)

// TemplateRenderer ...
type TemplateRenderer struct {
	templates map[string]*template.Template
}

// Render ...
func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tr.templates[name].ExecuteTemplate(w, "base.html", data)
}

func newTemplateRenderer() echo.Renderer {
	templates := make(map[string]*template.Template)
	templates["index"] = template.Must(template.ParseFiles("template/base.html", "template/index.html"))
	return &TemplateRenderer{templates: templates}
}

func main() {
	e := echo.New()
	e.Renderer = newTemplateRenderer()

	e.GET("/", handler.Index)
	e.GET("/api/hello", handler.Hello)

	e.Logger.Fatal(e.Start(":3000"))
}
