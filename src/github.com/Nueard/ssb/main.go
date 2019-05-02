package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/Nueard/ssb/content"
	"github.com/Nueard/ssb/context/page"
	"github.com/Nueard/ssb/context/site"
	"github.com/Nueard/ssb/render"
)

var tpl *template.Template
var siteResolver *site.Resolver
var pageResolver *page.Resolver
var contentLoader *content.Loader
var renderer *render.Renderer

func main() {
	domainMapLoader := site.NewDomainMapLoader()
	siteResolver = site.NewResolver(domainMapLoader)

	urlMapLoader := page.NewURLMapLoader()
	pageResolver = page.NewResolver(urlMapLoader)

	contentLoader = content.NewLoader()

	templateLoader := render.NewTemplateLoader("templates/")
	renderer = render.NewRenderer(templateLoader)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL)
	u, err := url.Parse(fmt.Sprintf("http://%s%s", r.Host, r.RequestURI))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}
	siteContext, err := siteResolver.Resolve(u)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	pageContext, err := pageResolver.Resolve(siteContext, u)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	content := contentLoader.Load(siteContext, pageContext)

	html, err := renderer.RenderHTML(content)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(html)
}
