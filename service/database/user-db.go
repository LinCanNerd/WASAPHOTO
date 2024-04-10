package database

import "database/sql"

func (db *appdbimpl) CreateUser(username string) (dbUser User, err error) {

	query := `INSERT INTO users (ID, username ) VALUES (?,?);`
	_, err = db.c.Exec(query, username, username)
	if err != nil {
		return
	}
	dbUser.Username = username
	dbUser.ID = username
	return dbUser, nil
}

func (db *appdbimpl) UpdateUsername(user User) (err error) {

	query := `UPDATE users SET username = ? WHERE ID = ?;`
	_, err = db.c.Exec(query, user.Username, user.ID)
	if err != nil {
		return err
	}

	query = "UPDATE comments SET authorUsername = ? WHERE authorID = ?;"
	_, err = db.c.Exec(query, user.Username, user.ID)
	if err != nil {
		return err
	}

	query = "UPDATE photos SET username = ? WHERE ownerID = ?;"
	_, err = db.c.Exec(query, user.Username, user.ID)
	if err != nil {
		return err
	}

	return
}

func (db *appdbimpl) GetUserID(username string) (user User, present bool, err error) {

	query := `SELECT * FROM users WHERE ID = ?;`
	err = db.c.QueryRow(query, username).Scan(&user.ID, &user.Username)
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

func (db *appdbimpl) GetUsername(usernameToSearch string) (user User, present bool, err error) {

	query := `SELECT * FROM users WHERE username = ?;`
	err = db.c.QueryRow(query, usernameToSearch).Scan(&user)
	if err != nil && err != sql.ErrNoRows {
		return user, true, err
	}
	return user, false, nil

}
