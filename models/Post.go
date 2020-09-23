package models

type Post struct {
	Id      string `db:"id"`
	Title   string `db:"title"`
	ContentHtml string `db:"content_html"`
	ContentMarkdown string `db:"content_markdown"`
}

func NewPost(id, title, contentHtml, contentMarkdown string) *Post {
	return &Post{id, title, contentHtml, contentMarkdown}
}
