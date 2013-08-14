package shawty

// import (
// 	"html/template"
// 	"net/http"
// )
// 
// type SiteConfig struct {
// 	StaticPath string
// }
// 
// var siteConfig = &SiteConfig{
// 	StaticPath:   "/static/",
// 	TemplatesDir: "templates", // TODO: Use this
// }
// 
// // Handler for the main page.
// func handleIndex(w http.ResponseWriter, r *http.Request, fromLookup bool) {
// 	// TODO: Handle redirect from failed lookup
// 	tpl := template.Must(template.ParseFiles("templates/index.html"))
// 	err := tpl.ExecuteTemplate(w, "index.html", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
// 
// // Handler for a shortcode
// func handleCode(w http.ResponseWriter, r *http.Request, code string) {
// 	url := lookupCode(code)
// 	if url == nil {
// 		// Code doesn't exist, redirect to the homepage (indicating we came form
// 		// a nonexistant lookup
// 		handleIndex(w, r, true)
// 	} else {
// 		// Code exists - redirect appropritely
// 		doRedirect(w, r, url)
// 	}
// }
// 
// // Lookup a code
// func lookupCode(code string) string {
// 	// TODO
// }
// 
// func doRedirect(w http.ResponseWriter, r *http.Request, url string) {
// 	// TODO: Some kind of logging
// 	http.Redirect(w, r, url, http.StatusFound)
// }
// 
// // Handle the attempted creation of a new code
// func handleNewCode(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	url := params["url"]
// 	code := params["code"]
// 	if len(url) == 0 {
// 		handleBadRequest(w, r)
// 	} else if len(code) == 0 {
// 		handleCreateCode(w, r, url[0])
// 	} else {
// 		handleCreateCustomCode(w, r, url[0], code[0])
// 	}
// }
// 
// func handleBadRequest(w http.ResponseWriter, r *http.Request) {
// 	// TODO
// }
// 
// func handleCreateCode(w http.ResponseWriter, r *http.Request, url string) {
// 	code := genNewCodeForUrl(url)
// 	handleCreateCustomCode(w, r, url, code)
// }
// 
// func handleCreateCustomCode(w http.ResponseWriter, r *http.Request, url string,
// 	code string) {
// 	
// }
// 
// func genNewCodeForUrl(url string) string {
// 
// }
// 
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Domain root: serve the homepage on GET, create shortened URLs on POST
// 		if code := r.URL.Path; code == "" {
// 			if r.Mehod == "GET" {
// 				handleIndex(w, r, false)
// 			} else {
// 				handleNewCode(w, r)
// 			}
// 		} else {
// 			handleCode(w, r, code)
// 		}
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
// 
// 
// 
// 
// 
// 

