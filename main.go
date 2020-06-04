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
	templates["hello"] = template.Must(template.ParseFiles("template/base.html", "template/hello.html"))
	templates["hello_form"] = template.Must(template.ParseFiles("template/base.html", "template/hello_form.html"))
	return &TemplateRenderer{templates: templates}
}

func main() {
	e := echo.New()
	e.Renderer = newTemplateRenderer()

	e.GET("/", handler.IndexGet)
	e.GET("/hello", handler.HelloGet)
	e.POST("/hello", handler.HelloPost)
	e.GET("/hello_form", handler.HelloFormGet)
	e.GET("/api/hello", handler.APIHelloGet)
	e.POST("/api/hello", handler.APIHelloPost)

	e.Logger.Fatal(e.Start(":3000"))
}
