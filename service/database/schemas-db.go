package database

import "time"

type User struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
}

type FullUser struct {
	ID        string  `json:"user_id"`
	Username  string  `json:"username"`
	Followers []User  `json:"followerList"`
	Following []User  `json:"followingList"`
	Banned    []User  `json:"bannedList"`
	Photo     []Photo `json:"photoList"`
}

type Photo struct {
	Owner    string    `json:"user_id"`
	Username string    `json:"username"`
	ID       int64     `json:"pid"`
	Likes    []User    `json:"likes"`
	Comments []Comment `json:"comments"`
	Date     time.Time `json:"date"`
	Format   string    `json:"format"`
}

type Comment struct {
	AuthorID       string    `json:"user_id"`
	AuthorUsername string    `json:"username"`
	PhotoID        int64     `json:"pid"`
	CommentId      int64     `json:"comment_id"`
	Text           string    `json:"text"`
	Date           time.Time `json:"date"`
}
