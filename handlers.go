package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goBlogExample/connection"
	"goBlogExample/models"
	"goBlogExample/utils"
	"html/template"
	"net/http"
	"time"
)

func getHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")

	html := utils.ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func indexHandler(rnd render.Render, r *http.Request) {
	cookie, err := r.Cookie("sessionId")
	if err == nil {
		session, e := inMemorySession.Get(cookie.Value)
		if e == nil {
			fmt.Println(session)
		}
	}

	posts, _ := connection.GetPosts()
	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	post, err := connection.ShowPost(id)
	if err != nil {
		http.NotFound(w, r)
	}

	rnd.HTML(200, "write", post)
}

func savePostHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
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

func deleteHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
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

func getLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "login", nil)
}

func postLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	//password := r.FormValue("password")

	sessionId := inMemorySession.Init(username)

	cookie := &http.Cookie{
		Name: "sessionId",
		Value: sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}