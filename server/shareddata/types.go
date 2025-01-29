package shareddata

import "time"

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Uuid      string `json:"uid"`
}

type Comment struct {
	UserId  int
	UserUID string
	PostId  int    `json:"postId"`
	Content string `json:"content"`
}

type ShowComment struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
	Date       string `json:"date"`
	IsLiked    bool   `json:"isliked"`
	IsDisliked bool   `json:"isdisliked"`
}
type Post struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Author        string    `json:"author"`
	Created       time.Time `json:"date"`
	Likes         int       `json:"likes"`
	Dislikes      int       `json:"dislikes"`
	IsLiked       bool      `json:"isliked"`
	IsDisliked    bool      `json:"isdisliked"`
	CommentsCount int       `json:"commentsCount"`
	Categories    []string  `json:"categories"`
	Joined_at     time.Time `json:"joined_at"`
}

type PostMetaData struct {
	PostsCount    int
	PostsPages    int
	StandardCount int
}

type CommentMetaData struct {
	CommentsCount int
	CommentsPages int
	StandardCount int
}

type InfoData struct {
	Authorize  bool     `json:"authorize"`
	Username   string   `json:"username"`
	Categories []string `json:"categories"`
}

type React struct {
	Thread_type string `json:"thread_type"`
	Thread_id   int    `json:"thread_id"`
	React       int    `json:"react"`
}

type ReactResponse struct {
	Like       int  `json:"Like"`
	Dislike    int  `json:"Dislike"`
	IsLiked    bool `json:"isliked"`
	IsDisliked bool `json:"isdisliked"`
}
