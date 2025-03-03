package main

import (
	"html/template"
	"net/http"
)

type App struct {
	templates map[string]*template.Template
}

func (app *App) Init() {
	app.templates = make(map[string]*template.Template)
}

func (app *App) ParseTemplates() {
	home, err := template.ParseFiles(
		"./templates/root.template",
		"./templates/home.template",
	)
	if err != nil {
		panic(err)
	}
	app.templates["home"] = home
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	app.templates["home"].ExecuteTemplate(w, "root.template", nil)
}

func main() {
	app := &App{}
	app.Init()
	app.ParseTemplates()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.Home)
	mux.HandleFunc("GET /index.html", app.Home)
	mux.Handle("GET /static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", mux)
}
