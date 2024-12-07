package utils

type User struct {
	Id       int
	Username string
	Email    string
	Password string
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
	Id int
	UserId int
	Title string
	Content string
	Image string
	Categories string
	Date string
}

type Reaction struct {
	Id int
	UserId int
	PostId int
	CommentId int
	Type string
}

type PostCategory struct {
	PostId int
	CategoryId int
}