package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/session"
	"html/template"
)

var inMemorySession *session.Session

func main() {
	connection.Connect()
	defer connection.CloseConnection()

	inMemorySession = session.NewSession()

	runServer()
}

func runServer() {
	fmt.Println("Listen port :3000")

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

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