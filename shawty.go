package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := &Page{Title: "Shawty"}
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		err := tpl.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
