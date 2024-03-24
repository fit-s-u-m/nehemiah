package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "home.html", nil)
	})
	e.GET("/blogs", func(c echo.Context) error {
		return c.Render(200, "blog.html", nil)
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
