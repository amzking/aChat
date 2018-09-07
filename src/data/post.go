package data

import "time"

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (post *Post) CreateAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2018 at 10.58pm")
}
