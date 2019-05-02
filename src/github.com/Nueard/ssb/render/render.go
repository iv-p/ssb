package render

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Nueard/ssb/content"
)

// Renderer gets a page content and generates its html
type Renderer struct {
	tpl *template.Template
}

// NewRenderer creates and returns a new renderer
func NewRenderer(templateLoader *TemplateLoader) *Renderer {
	return &Renderer{
		tpl: templateLoader.Load(),
	}
}

// RenderHTML takes a page content object and returns the generated HTML for that page
func (r *Renderer) RenderHTML(c content.Content) ([]byte, error) {
	layout := r.tpl.Lookup(c.Template)
	if layout == nil {
		return []byte{}, fmt.Errorf("template %s not found", c.Template)
	}
	w := bytes.NewBuffer([]byte{})
	err := layout.Execute(w, c)
	return w.Bytes(), err
}
