package handlers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/session"
)

func IndexHandler(rnd render.Render, s *session.Session) {
	fmt.Println(s, s.Username, "session")

	posts, _ := connection.GetPosts()
	rnd.HTML(200, "index", posts)
}