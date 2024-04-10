package database

func (db *appdbimpl) LikePhoto(photoID int64, userID string) (err error) {

	_, err = db.c.Exec("INSERT INTO likes (photoID, likerID) VALUES (?,?);", photoID, userID)
	if err != nil {
		return
	}
	return
}

func (db *appdbimpl) UnlikePhoto(photoID int64, userID string) (err error) {

	_, err = db.c.Exec("DELETE FROM likes WHERE photoID = ? AND likerID = ?;", photoID, userID)
	if err != nil {
		return
	}
	return
}
