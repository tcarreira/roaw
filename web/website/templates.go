package website

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/config"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewRenderer() *TemplateRenderer {
	t := TemplateRenderer{
		Templates: template.Must(
			template.ParseGlob(filepath.Join(config.WebsiteTemplatesPath(), "*.html.j2")),
		),
	}
	return &t
}
