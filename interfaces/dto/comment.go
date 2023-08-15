package dto

import "time"

type CommentsItem struct {
	Id         int64     `json:"id"`
	Pid        int64     `json:"pid"`
	ArticleId  int64     `json:"article_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Url        string    `json:"url"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

type CommentCount struct {
	Count int `json:"count"`
}

type AddComment struct {
	LastId    int    `json:"last_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Url       string `json:"url"`
	Content   string `json:"content"`
	ArticleId int    `json:"article_id"`
	Pid       int    `json:"pid"`
}
