package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"goBlogExample/models"
	"log"
)

var dbConnection *sqlx.DB

var schema = `
CREATE TABLE posts (
    id varchar(16),
    title varchar(255),
    content_html text,
    content_markdown text
);
`

func connect() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=go_blog_example password=123123 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	dbConnection = db

	//db.MustExec(schema)
}

func closeConnection() {
	dbConnection.Close()
}

func getPosts() (posts []models.Post, err error) {
	err = dbConnection.Select(&posts, "SELECT * FROM Posts")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func showPost(id string) (post models.Post, err error) {
	err = dbConnection.Get(&post, "SELECT * FROM Posts WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func updatePost(post models.Post) (err error) {
	_, err = dbConnection.Exec(
		"UPDATE Posts SET title = $1, content_html = $2, content_markdown = $3 WHERE id = $4",
		post.Title,
		post.ContentHtml,
		post.ContentMarkdown,
		post.Id,
	)
	return
}

func createPost(post models.Post) (err error) {
	_, err = dbConnection.Exec(
		"INSERT INTO Posts (id, title, content_html, content_markdown) VALUES ($1, $2, $3, $4)",
		post.Id,
		post.Title,
		post.ContentHtml,
		post.ContentMarkdown,
	)
	return
}

func deletePost(post models.Post) (err error) {
	_, err = dbConnection.Exec(
		"DELETE FROM Posts WHERE id = $1",
		post.Id,
	)
	return
}
