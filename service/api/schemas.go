package api

import (
	"time"

	"github.com/LinCannn/WASA/service/database"
)

type User struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
}

type FullUser struct {
	ID        string           `json:"user_id"`
	Username  string           `json:"username"`
	Followers []database.User  `json:"followerList"`
	Following []database.User  `json:"followingList"`
	Banned    []database.User  `json:"bannedList"`
	Photo     []database.Photo `json:"photoList"`
}

type Photo struct {
	Owner    string             `json:"user_id"`
	Username string             `json:"username"`
	ID       int64              `json:"pid"`
	Likes    []database.User    `json:"likes"`
	Comments []database.Comment `json:"comments"`
	Date     time.Time          `json:"date"`
	Format   string             `json:"format"`
}

type Comment struct {
	AuthorID       string    `json:"user_id"`
	AuthorUsername string    `json:"username"`
	PhotoID        int64     `json:"pid"`
	CommentId      int64     `json:"comment_id"`
	Text           string    `json:"text"`
	Date           time.Time `json:"date"`
}

// ToDatabase are functions that transform the api correspective to their database correspective
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:       u.ID,
		Username: u.Username,
	}
}

func (u *User) FromDatabase(dbUser database.User) {
	u.ID = dbUser.ID
	u.Username = dbUser.Username
}

func (fu *FullUser) ToDatabase() (dbFullUser database.FullUser) {
	return database.FullUser{
		ID:        fu.ID,
		Username:  fu.Username,
		Followers: fu.Followers,
		Following: fu.Following,
		Banned:    fu.Banned,
		Photo:     fu.Photo,
	}
}

func (fu *FullUser) FromDatabase(dbFullUser database.FullUser) {
	fu.ID = dbFullUser.ID
	fu.Username = dbFullUser.Username
	fu.Followers = dbFullUser.Followers
	fu.Following = dbFullUser.Following
	fu.Banned = dbFullUser.Banned
	fu.Photo = dbFullUser.Photo
}

// Photo

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		Owner:    p.Owner,
		Username: p.Username,
		ID:       p.ID,
		Likes:    p.Likes,
		Comments: p.Comments,
		Date:     p.Date,
		Format:   p.Format,
	}
}

func (p *Photo) FromDatabase(dbPhoto database.Photo) {
	p.Owner = dbPhoto.Owner
	p.Username = dbPhoto.Username
	p.ID = dbPhoto.ID
	p.Likes = dbPhoto.Likes
	p.Comments = dbPhoto.Comments
	p.Date = dbPhoto.Date
	p.Format = dbPhoto.Format
}

// Comment

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		AuthorID:       c.AuthorID,
		AuthorUsername: c.AuthorUsername,
		PhotoID:        c.PhotoID,
		CommentId:      c.CommentId,
		Text:           c.Text,
		Date:           c.Date,
	}
}

func (c *Comment) FromDatabase(dbComment database.Comment) {
	c.AuthorID = dbComment.AuthorID
	c.AuthorUsername = dbComment.AuthorUsername
	c.CommentId = dbComment.CommentId
	c.PhotoID = dbComment.PhotoID
	c.Text = dbComment.Text
	c.Date = dbComment.Date
}
