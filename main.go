package shawty

import (
	"fmt"
	"html/template"
	"net/http"
)

// Handle an error by throwing a 500
func handleErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler for the main page.
func handleIndex(w http.ResponseWriter, fromLookup bool) {
	// TODO: Handle redirect from failed lookup
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	handleErr(w, tpl.ExecuteTemplate(w, "index.html", nil))
}

// Render a normal JSON response
func jsonResponse(w http.ResponseWriter, r *http.Request, url string, code string, errTxt string) {
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Domain root: serve the homepage on GET, create shortened URLs on POST
		if r.Method == "GET" {
			if code := r.URL.Path; code == "" {
				handleIndex(w, false)
			} else {
				url, err := LookupUrl(code)
				handleErr(w, err)

				http.Redirect(w, r, url, http.StatusFound)
			}
		} else {
			// Extract and validate URL parameter
			params := r.URL.Query()
			url := params["url"][0]
			if len(url) == 1 {
				// TODO handle too many params
			}

			// TODO validate URL

			if code := r.URL.Path; code == "" {
				// POSTing to root: generate a random shortcode
				code, err := ShortenUrl(url)
				handleErr(w, err)

				jsonResponse(w, r, url, code, "")
			} else {
				// POSTing to a path: map to the given shortcode
				success, err := ShortenUrlToCode(url, code)
				handleErr(w, err)

				if success == false {
					jsonResponse(w, r, url, "", fmt.Sprintf("The code %s is taken."))
				}

				jsonResponse(w, r, url, code, "")
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
