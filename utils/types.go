package utils

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

type PostsResult struct {
	Id         int
	UserName   string
	Title      string
	Content    string
	Categories []string
	Date       string
	Reactions  Reactions
}
type CommentType struct {
	Id        int
	Username  string
	UserImage string
	Comment   string
	Date      string
	Reactions Reactions
}

type Reaction struct {
	Id        int
	UserId    int
	PostId    int
	CommentId int
	Type      string
}

type Reactions struct {
	Likes    int
	Dislikes int
	Action   string
}
type PostCategory struct {
	PostId     int
	CategoryId int
}

type Error struct {
	Err  error
	Code int
}
