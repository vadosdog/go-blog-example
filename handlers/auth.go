package handlers

import (
	"github.com/martini-contrib/render"
	"goBlogExample/session"

	"net/http"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	//password := r.FormValue("password")

	s.Username = username
	s.IsAuthorized = true

	rnd.Redirect("/")
}

func LogoutHandler(rnd render.Render, s *session.Session) {
	s.Username = ""
	s.IsAuthorized = false

	rnd.Redirect("/")
}