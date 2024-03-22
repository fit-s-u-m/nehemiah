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

type blog_list struct {	
	Title string
	Content string
}
bloglists := [3]blog_list{ 
	blog_list{
		Title: "Trigonometry",
		Content: "Trigonometry is a branch of mathematics that studies relationships between the sides and angles of triangles."
	}, 
	blog_list{
		Title: "Title 2", 
		Content: "Content 2"
	}, 
	blog_list{Title: "Title 3", Content: "Content 3"}
}



func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "home.html", nil)
	})
	e.POST("/blogs", func(c echo.Context) error {
		return c.Render(200, "blog.html", blog_lists)
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))

}
