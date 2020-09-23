package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
)


func main() {
	connect()
	defer closeConnection()

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

	staticOptions := martini.StaticOptions{Prefix: "assets"}

	m.Use(martini.Static("assets", staticOptions))
	m.Get("/", indexHandler)
	m.Get("/posts/new", writeHandler)
	m.Get("/posts/:id", editHandler)
	m.Post("/posts/save", savePostHandler)
	m.Get("/posts/:id/delete", deleteHandler)
	m.Post("/getHtml", getHtmlHandler)

	m.Run()
}