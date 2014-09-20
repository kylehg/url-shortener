package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

type App struct {
	params httprouter.Params
	req *http.Request
	res http.ResponseWriter
	router http.Handler
}

todo new app function

func (app App) route(path string, method string, handle func(*App)) {
	fmt.Printf("%s %s\n", method, path)
	handler := func (res http.ResponseWriter, req *http.Request) {
		app.res = res
		app.req = req
		handle(app)
	}
	if "GET" == method {
		app.router.GET(path, handler)
	} else if "POST" == method {
		app.router.POST(path, handler)
	}
}

func (app App) writeJson(data Json, statusCode int) {
	app.res.WriteHeader(statusCode)
	app.res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, data)
}

func (app App) writeJsonErr(err error) {
	app.writeJson(Json{"error": err.Error()}, http.StatusInternalServerError)
}

func (app App) redirect(url string) {
	http.Redirect(app.res, app.req, url, http.StatusFound)
}

func (app App) getFormParam(param string) string {
	if err := app.req.ParseForm(); err != nil {
		jsonErr(err, w)
		return ""
	}
	return app.req.Form.Get(param)
}

func (app App) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

// GET /
func handleIndex(app *App) {
	app.res.Write([]byte("Index"))
}

// POST /
func handleCreate(app *App) {
	url := app.getFormParam(URL_KEY)
	code, err := shawty.ShortenDefault(url)
	if err != nil {
		app.writeJsonErr(err)
		return
	}
	app.writeJson(Json{"code": code, "url": url}, http.StatusCreated)
}

// GET /:shortcode
func handleLookup(app *App) {
	code := app.req.URL.Path[1:]
	url, err := shawty.GetUrl(code)
	if err != nil {
		app.writeJsonErr(err, w)
		return
	}
	app.redirect(url)
}

// POST /:shortcode
func handleCreateWithCode(app *App) {
	code := app.req.URL.Path[1:]
	url := app.getFormParam(URL_KEY)
	if err := shawty.ShortenCustom(url, code); err != nil {
		app.writeJsonErr(err)
		return
	}
	app.writeJson(Json{"code": code, "url": url}, http.StatusCreated)
}

func main() {
	app := newApp()
	app.route("/", "GET", handleIndex)
	app.route("/", "POST", handleCreate)
	app.route("/:hash", "GET", handleLookup)
	app.route("/:hash", "POST", handleCreateWithCode)

	fmt.Printf("Server running on port 8080...\n")
	http.ListenAndServe(":8080", app)
}
