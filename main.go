package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/handlers"
	"goBlogExample/session"
	"html/template"
)


func main() {
	connection.Connect()
	defer connection.CloseConnection()

	session.StartSession()

	runServer()
}

func runServer() {
	fmt.Println("Listen port :3000")

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": handlers.Unescape}

	m.Use(session.Middleware)

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{unescapeFuncMap},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	routes(m)

	m.Run()
}