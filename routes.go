package main

import "github.com/go-martini/martini"

func routes(m *martini.ClassicMartini) {
	staticOptions := martini.StaticOptions{Prefix: "assets"}

	m.Use(martini.Static("assets", staticOptions))
	m.Get("/", indexHandler)
	m.Get("/posts/new", writeHandler)
	m.Get("/posts/:id", editHandler)
	m.Post("/posts/save", savePostHandler)
	m.Get("/posts/:id/delete", deleteHandler)
	m.Post("/getHtml", getHtmlHandler)
	m.Get("/login", getLoginHandler)
	m.Post("/login", postLoginHandler)
}