package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"

	fragmentloader "github.com/Nueard/ssb/content/fragment-loader"
)

// ITemplateLoader is the interface of TemplateLoader
type ITemplateLoader interface {
	Load(string) *template.Template
}

// TemplateLoader loads html template files
type TemplateLoader struct {
	dir string
}

// NewTemplateLoader creates and returns a new template loader
func NewTemplateLoader(dir string) *TemplateLoader {
	return &TemplateLoader{
		dir: dir,
	}
}

// Load traverses a directory and loads all the templates setting
// the names of the templates to the relative filepath to the directory
func (tl *TemplateLoader) Load() *template.Template {
	tpl := template.New("")

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"CallTemplate": func(fragment fragmentloader.Fragment) (ret template.HTML, err error) {
			buf := bytes.NewBuffer([]byte{})
			err = tpl.ExecuteTemplate(buf, fragment.Template, fragment)
			ret = template.HTML(buf.String())
			return
		},
	}

	files, err := filepath.Glob(fmt.Sprintf("%s**/*.html", tl.dir))
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		file, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}

		f = strings.Replace(f, tl.dir, "", -1)

		_, err = tpl.New(f).Funcs(funcMap).Parse(string(file))
		if err != nil {
			panic(err)
		}
	}

	return tpl
}
