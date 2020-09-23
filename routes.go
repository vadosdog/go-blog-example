package main

import (
	"github.com/go-martini/martini"
	"goBlogExample/handlers"
)

func routes(m *martini.ClassicMartini) {
	staticOptions := martini.StaticOptions{Prefix: "assets"}

	m.Use(martini.Static("assets", staticOptions))
	m.Get("/", handlers.IndexHandler)
	m.Get("/posts/new", handlers.WriteHandler)
	m.Get("/posts/:id", handlers.ViewHandler)
	m.Get("/posts/:id/edit", handlers.EditHandler)
	m.Post("/posts/save", handlers.SavePostHandler)
	m.Get("/posts/:id/delete", handlers.DeleteHandler)
	m.Post("/getHtml", handlers.GetHtmlHandler)
	m.Get("/login", handlers.GetLoginHandler)
	m.Post("/login", handlers.PostLoginHandler)
	m.Get("/logout", handlers.LogoutHandler)
}