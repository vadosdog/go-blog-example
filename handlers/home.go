package handlers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/models"
	"goBlogExample/session"
)

func IndexHandler(rnd render.Render, s *session.Session) {
	fmt.Println(s, s.Username, "session")

	posts, _ := connection.GetPosts()

	model := models.PostListModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Posts = posts

	rnd.HTML(200, "index", model)
}