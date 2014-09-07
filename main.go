package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kylehg/shawty/shawty"
)

const (
	SHORTCODE_KEY = "code"
	URL_KEY       = "url"
)

type Json map[string]interface{}

func (j Json) String() string {
	bytes, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func writeJson(data Json, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, data)
}

func jsonErrCode(err error, errCode int, w http.ResponseWriter) {
	w.WriteHeader(errCode)
	data := Json{"error": err.Error()}
	writeJson(data, w)
}

func jsonErr(err error, w http.ResponseWriter) {
	jsonErrCode(err, http.StatusInternalServerError, w)
}

func jsonRes(data Json, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	writeJson(data, w)
}

func getFormParam(param string, w http.ResponseWriter, r *http.Request) string {
	if err := r.ParseForm(); err != nil {
		jsonErr(err, w)
		return ""
	}
	return r.Form.Get(param)
}

// GET /
func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}

// POST /
func handleCreate(w http.ResponseWriter, r *http.Request) {
	url := getFormParam(URL_KEY, w, r)
	code, err := shawty.ShortenDefault(url)
	if err != nil {
		jsonErr(err, w)
		return
	}
	jsonRes(Json{"code": code, "url": url}, http.StatusCreated, w)
}

// GET /:shortcode
func handleLookup(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	url, err := shawty.GetUrl(code)
	if err != nil {
		jsonErr(err, w)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// POST /:shortcode
func handleCreateWithCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	url := getFormParam(URL_KEY, w, r)
	if err := shawty.ShortenCustom(url, code); err != nil {
		jsonErr(err, w)
		return
	}
	jsonRes(Json{"code": code, "url": url}, http.StatusCreated, w)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Domain root: serve the homepage on GET, create shortened URLs on POST
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		if r.URL.Path == "/" {
			switch r.Method {
			case "GET":
				handleIndex(w, r)
			case "POST":
				handleCreate(w, r)
			}
		} else {
			switch r.Method {
			case "GET":
				handleLookup(w, r)
			case "POST":
				handleCreateWithCode(w, r)
			}
		}
	})

	fmt.Printf("Server running on port 8080...\n")
	http.ListenAndServe(":8080", nil)
}
