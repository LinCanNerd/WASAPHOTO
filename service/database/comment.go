package database

import "database/sql"

func (db *appdbimpl) CommentPhoto(comment Comment) (dbComment Comment, err error) {
	query, err := db.c.Exec("INSERT INTO comments (photoID, authorID, authorUsername, text, date) VALUES (?,?,?,?,?);", comment.PhotoID, comment.AuthorID, comment.AuthorUsername, comment.Text, comment.Date)
	if err != nil {
		return
	}
	dbComment = comment
	commentID, err := query.LastInsertId()
	dbComment.CommentId = commentID
	return
}

func (db *appdbimpl) UncommentPhoto(commentID int64) (err error) {
	_, err = db.c.Exec("DELETE FROM comments WHERE commentID = ?;", commentID)
	if err != nil {
		return
	}
	return
}

func (db *appdbimpl) GetCommentByID(commentID int64) (dbComment Comment, present bool, err error) {
	query := "SELECT * FROM comments WHERE commentID = ?;"
	row := db.c.QueryRow(query, commentID)
	err = row.Scan(&dbComment.CommentId, &dbComment.PhotoID, &dbComment.AuthorID, &dbComment.AuthorUsername, &dbComment.Text, &dbComment.Date)
	if err != nil && err != sql.ErrNoRows {
		return
	} else if err == sql.ErrNoRows {
		err = nil
		return

	} else {
		err = nil
		present = true
		return
	}
}
