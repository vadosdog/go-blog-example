package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goBlogExample/models"
	"html/template"
	"net/http"
)

func getHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")

	html := ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func indexHandler(rnd render.Render) {
	posts, _ := getPosts()
	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	post, err := showPost(id)
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
	contentHtml := ConvertMarkdownToHtml(contentMarkdown)
	newItem := id == ""

	if newItem {
		post = models.Post{Id: GenerateId(), Title: title, ContentHtml: contentHtml, ContentMarkdown: contentMarkdown}
		err = createPost(post)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		post, err = showPost(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
		err = updatePost(post)
		if err != nil {
			fmt.Println(err)
		}
	}

	rnd.Redirect("/")
}

func deleteHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]

	post, err := showPost(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = deletePost(post)
	if err != nil {
		fmt.Println(err)
	}

	rnd.Redirect("/")
}
