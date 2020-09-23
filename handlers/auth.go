package handlers

import (
	"github.com/martini-contrib/render"
	"goBlogExample/session"

	"net/http"
)

func GetLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	//password := r.FormValue("password")

	s.Username = username

	rnd.Redirect("/")
}