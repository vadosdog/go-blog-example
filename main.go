package main

import (
	"fmt"
	"goBlogExample/models"
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
)

var posts map[string]*models.Post

func main() {
	fmt.Println("Listen port :3000")

	posts = make(map[string]*models.Post, 0)

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{unescapeFuncMap},
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

func getHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")

	htmlBytes := blackfriday.MarkdownBasic([]byte(md))

	rnd.JSON(200, map[string]interface{} {"html": string(htmlBytes)})
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func indexHandler(rnd render.Render) {
	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	rnd.HTML(200, "write", post)
}

func savePostHandler(rnd render.Render, r *http.Request) {
	var id string
	var post *models.Post

	id = r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := string(blackfriday.MarkdownBasic([]byte(contentMarkdown)))

	if id != "" {
		post = posts[id]
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
	} else {
		id = GenerateId()
		post = models.NewPost(id, title, contentHtml, contentMarkdown)
		posts[post.Id] = post
	}

	rnd.Redirect("/")
}

func deleteHandler(rnd render.Render, r *http.Request, params martini.Params) {
	id := params["id"]

	delete(posts, id)

	rnd.Redirect("/")
}
