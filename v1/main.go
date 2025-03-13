package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HttpError struct {
	Status int
	Err    error
}

func NewHttpError(status int, err error) *HttpError {
	return &HttpError{Status: status, Err: err}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d %s", e.Status, http.StatusText(e.Status))
}

func (e *HttpError) Unwrap() error {
	return e.Err
}

type AppHandler func(http.ResponseWriter, *http.Request) error

func (f AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		switch v := err.(type) {
		case *HttpError:
			http.Error(w, v.Error(), v.Status)
		default:
			http.Error(w, v.Error(), 500)
		}
	}
}

type App struct {
	templates map[string]*template.Template
}

func (app *App) InitTemplates() {
	app.templates = make(map[string]*template.Template)
	home, err := template.ParseFiles(
		"./templates/root.template",
		"./templates/home.template",
	)
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	app.templates["home"] = home
}

func (app *App) Init() {
	app.InitTemplates()
}

func (app *App) Home() AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		buf := new(bytes.Buffer)
		// tpl.render(w, "home")
		tpl := app.templates["home"]
		err := tpl.ExecuteTemplate(buf, "root!.template", nil)
		if err != nil {
			return NewHttpError(500, err)
		}
		buf.WriteTo(w)
		return nil
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	app := &App{}
	app.Init()
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", app.Home())
	mux.Handle("GET /index.html", app.Home())
	mux.Handle("GET /static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", mux)
}
