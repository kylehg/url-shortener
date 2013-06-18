package main

import (
	"github.com/garyburd/redigo/redis"
	"html/template"
	"net/http"
)

type SiteConfig struct {
	StaticPath string
}

var siteConfig = &SiteConfig{
	StaticPath: "/static/",
	TemplatesDir: "templates" // TODO: Use this
}


// Handler for the main page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tpl.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


// Handler for a shortcode
func codeHandler(w http.ResponseWriter, r *http.Request, code string) {
	
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if code := r.URL.Path; code == "" {
			indexHandler(w, r)
		} else {
			codeHandler(w, r, code)
		}
	})
	http.ListenAndServe(":8080", nil)
}
