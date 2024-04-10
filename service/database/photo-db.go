package database

import (
	"database/sql"
)

func (db *appdbimpl) CreatePhoto(photo Photo) (dbPhoto Photo, err error) {

	query, err := db.c.Exec("INSERT INTO photos (ownerID, username, date, format) VALUES (?,?,?,?);", photo.Owner, photo.Username, photo.Date, photo.Format)
	if err != nil {
		return
	}
	dbPhoto = photo
	photoID, err := query.LastInsertId()
	dbPhoto.ID = photoID
	return
}

func (db *appdbimpl) DeletePhoto(ID int64) (err error) {

	_, err = db.c.Exec("DELETE FROM photos WHERE photoID= ?;", ID)
	if err != nil {
		return
	}
	return

}

func (db *appdbimpl) GetPhoto(ID int64) (dbPhoto Photo, present bool, err error) {

	query := "SELECT * FROM photos WHERE photoID = ?;"

	row := db.c.QueryRow(query, ID)
	err = row.Scan(&dbPhoto.ID, &dbPhoto.Owner, &dbPhoto.Username, &dbPhoto.Format, &dbPhoto.Date)
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
