package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"os"
)

type TemplateRenderer struct {
	templates *template.Template
}

func InitRender(e *echo.Echo) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(wd + "/template/*/*.html")),
	}
	e.Renderer = renderer

	return nil
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
