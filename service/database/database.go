/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	Ping() error

	// user
	GetUserID(username string) (User, bool, error)
	GetUsername(usernameToSearch string) (User, bool, error)
	UpdateUsername(user User) error
	SearchUser(usernameToSearch string) (usersList []User, err error)
	CreateUser(username string) (User, error)

	// follow
	FollowUser(followerID string, followedID string) error
	UnfollowUser(followerID string, followedID string) error
	CheckFollow(followerID string, followedID string) (bool, error)
	RemoveFollowBothDirections(user1ID string, user2ID string) error

	// ban
	BanUser(banner User, banned User) error
	UnbanUser(banner User, banned User) error
	BannedUserCheck(requestingUser User, targetUser User) (bool, error)
	NoMoreFriends(banner User, banned User) (err error)

	// photo
	CreatePhoto(photo Photo) (Photo, error)
	DeletePhoto(ID int64) error
	GetPhoto(ID int64) (Photo, bool, error)

	// like
	LikePhoto(photoID int64, likerID string) error
	UnlikePhoto(photoID int64, likerID string) error

	// comment
	CommentPhoto(comment Comment) (Comment, error)
	UncommentPhoto(ID int64) error
	GetCommentByID(ID int64) (Comment, bool, error)

	// get
	GetFollowers(user User) ([]User, error)
	GetFollowing(user User) ([]User, error)
	GetBannedList(user User) ([]User, error)
	GetPhotosList(targetUser User) ([]Photo, error)
	GetCommentsList(photo Photo) ([]Comment, error)
	GetLikesList(photo int64) ([]User, error)
	GetMyStream(requestingID string) ([]Photo, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func createDatabase(db *sql.DB) error {
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS users (
			ID VARCHAR(16) NOT NULL PRIMARY KEY,
			username VARCHAR(16) NOT NULL
			);`,

		`CREATE TABLE IF NOT EXISTS following (
			followerID VARCHAR(16) NOT NULL REFERENCES users(ID),
			followedID VARCHAR(16) NOT NULL REFERENCES users(ID),
			PRIMARY KEY (followerID, followedID)
		);`,

		`CREATE TABLE IF NOT EXISTS banned (
			bannerID VARCHAR(16) NOT NULL REFERENCES users(ID),
			bannedID VARCHAR(16) NOT NULL REFERENCES users(ID),
			PRIMARY KEY (bannerID, bannedID)
		);`,

		`CREATE TABLE IF NOT EXISTS photos (
			photoID INTEGER PRIMARY KEY AUTOINCREMENT,
			ownerID VARCHAR(16) NOT NULL REFERENCES users(ID),
			username VARCHAR(16) NOT NULL REFERENCES users(username),
			format VARCHAR(3) NOT NULL,
			date DATETIME NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS likes (
			photoID INTEGER NOT NULL REFERENCES photos(photoID),
			likerID VARCHAR(16) NOT NULL REFERENCES users(ID),
			PRIMARY KEY (photoID, likerID)
		);`,

		`CREATE TABLE IF NOT EXISTS comments (
			commentID INTEGER PRIMARY KEY AUTOINCREMENT,
			photoID INTEGER NOT NULL REFERENCES photos(photoID),
			authorID VARCHAR(16) NOT NULL REFERENCES users(ID),
			authorUsername VARCHAR(16) NOT NULL REFERENCES users(username),
			text TEXT NOT NULL,
			date DATETIME NOT NULL
		);`,
	}
	for t := 0; t < len(tables); t++ {
		sqlStmt := tables[t]
		_, err := db.Exec(sqlStmt)

		if err != nil {
			return err
		}
	}

	return nil
}
