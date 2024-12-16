package utils

import "net/http"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfPass string `json:"password_config"`
}

type Session struct {
	Id     int
	UserId int
	Token  string
	Date   string
}

type Category struct {
	Id   int
	Name string
}

type Comment struct {
	Id      int
	UserId  int
	PostId  int
	Comment string
	Date    string
}

type Post struct {
	Id         int
	UserId     int
	Username   string
	Title      string
	Content    string
	Image      string
	Categories string
	Date       string
}

type Reaction struct {
	Id        int
	UserId    int
	PostId    int
	CommentId int
	Type      string
}

type PostCategory struct {
	PostId     int
	CategoryId int
}

type Error struct {
	W       http.ResponseWriter
	Err     error
	Code    int
}
