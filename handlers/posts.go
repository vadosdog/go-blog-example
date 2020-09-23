package handlers

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/models"
	"goBlogExample/utils"
	"html/template"
	"net/http"
)

func WriteHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func EditHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	post, err := connection.ShowPost(id)
	if err != nil {
		http.NotFound(w, r)
	}

	rnd.HTML(200, "write", post)
}

func SavePostHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	var id string
	var post models.Post
	var err error

	id = r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)
	newItem := id == ""

	if newItem {
		post = models.Post{Id: utils.GenerateId(), Title: title, ContentHtml: contentHtml, ContentMarkdown: contentMarkdown}
		err = connection.CreatePost(post)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		post, err = connection.ShowPost(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
		err = connection.UpdatePost(post)
		if err != nil {
			fmt.Println(err)
		}
	}

	rnd.Redirect("/")
}

func DeleteHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]

	post, err := connection.ShowPost(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = connection.DeletePost(post)
	if err != nil {
		fmt.Println(err)
	}

	rnd.Redirect("/")
}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")

	html := utils.ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}

func Unescape(x string) interface{} {
	return template.HTML(x)
}