package domain

type PostId int

type Post struct {
	ID      PostId
	Title   string
	Content string
	Author  string
}
