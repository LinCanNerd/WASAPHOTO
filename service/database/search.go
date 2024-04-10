package database

func (db *appdbimpl) SearchUser(usernameToSearch string) (usersList []User, err error) {

	query := "SELECT * FROM users WHERE username LIKE ?"

	rows, err := db.c.Query(query, usernameToSearch+"%")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return
		}
		usersList = append(usersList, user)
	}

	err = rows.Err()
	return
}
