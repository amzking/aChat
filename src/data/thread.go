package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3.04pm")
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("Select count(*) from posts where thread_id=$1", thread.Id)
	if err != nil {
		log.Fatal("no data")
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
