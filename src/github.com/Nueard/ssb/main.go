package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Nueard/ssb/site"
)

var tpl *template.Template
var sites = map[string]string{
	"dummydomain.com":           "site1",
	"subdomain.dummydomain.com": "site2",
}
var resolver *site.Resolver

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	resolver = site.NewResolver(sites)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	path := ("home.html")
	siteID, err := resolver.Resolve(r.Host)
	if err != nil {
		log.Print(err)
	}

	log.Print(siteID)
	layout := tpl.Lookup("layout.html")
	layout, _ = layout.Clone()
	t := tpl.Lookup(path)
	_, _ = layout.AddParseTree("content", t.Tree)
	layout.Execute(w, "")
}
