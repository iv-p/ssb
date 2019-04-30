package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	path := ("home.html")

	layout := tpl.Lookup("layout.html")
	layout, _ = layout.Clone()
	t := tpl.Lookup(path)
	_, _ = layout.AddParseTree("content", t.Tree)
	layout.Execute(w, "")
}
