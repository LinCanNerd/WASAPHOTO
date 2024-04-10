package database

// Database function che recupera la lista degli utenti che seguono l'utente specificato
func (db *appdbimpl) GetFollowers(requestingUser User) ([]User, error) {
	// Utilizza una query SQL per selezionare tutti gli utenti che seguono l'utente specificato.
	rows, err := db.c.Query("SELECT followerID FROM following WHERE followedID = ?", requestingUser.ID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that follow the requesting user)
	var followers []User
	for rows.Next() {
		var follower User
		err = rows.Scan(&follower.ID)
		if err != nil {
			return nil, err
		}
		follower, present, err := db.GetUserID(follower.ID)
		if err != nil || !present {
			return nil, err
		}

		followers = append(followers, follower)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}

func (db *appdbimpl) GetFollowing(requestingUser User) ([]User, error) {
	// Utilizza una query SQL per selezionare tutti gli utenti che l'utente specificato segue.
	rows, err := db.c.Query("SELECT followedID FROM following WHERE followerID = ?", requestingUser.ID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that the requesting user follows)
	var following []User
	for rows.Next() {
		var followed User
		err = rows.Scan(&followed.ID)
		if err != nil {
			return nil, err
		}
		followed, present, err := db.GetUserID(followed.ID)
		if err != nil || !present {
			return nil, err
		}
		following = append(following, followed)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}

func (db *appdbimpl) GetBannedList(requestingUser User) ([]User, error) {
	// Utilizza una query SQL per selezionare tutti gli utenti che l'utente specificato ha bannato.
	rows, err := db.c.Query("SELECT bannedID FROM banned WHERE bannerID = ?", requestingUser.ID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that the requesting user follows)
	var banned []User
	for rows.Next() {
		var bannedUser User
		err = rows.Scan(&bannedUser.ID)
		if err != nil {
			return nil, err
		}
		bannedUser, present, err := db.GetUserID(bannedUser.ID)
		if err != nil || !present {
			return nil, err
		}
		banned = append(banned, bannedUser)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return banned, nil
}

func (db *appdbimpl) GetPhotosList(targetUser User) ([]Photo, error) { // requestinUser User,
	// esegue una query SQL per selezionare tutte le foto di targetUser e le ordina in base alla data in ordine decrescente.
	rows, err := db.c.Query("SELECT * FROM photos WHERE ownerID = ? ORDER BY date DESC", targetUser.ID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()
	var photos []Photo
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.Owner, &photo.Username, &photo.Format, &photo.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}

func (db *appdbimpl) GetCommentsList(photo Photo) ([]Comment, error) {
	// Esegue una query SQL per selezionare tutti i commenti associati a una foto specificata.
	rows, err := db.c.Query("SELECT * FROM comments WHERE photoID = ?",
		photo.ID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the comments in the resulset
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.CommentId, &comment.PhotoID, &comment.AuthorID, &comment.AuthorUsername, &comment.Text, &comment.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}

func (db *appdbimpl) GetLikesList(photoID int64) ([]User, error) {
	// Esegue una query SQL per selezionare tutti gli utenti che hanno messo "mi piace" a una foto specificata.
	rows, err := db.c.Query("SELECT likerID FROM likes WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset
	var likes []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}

func (db *appdbimpl) GetMyStream(requestingID string) ([]Photo, error) {

	rows, err := db.c.Query("SELECT photoID, ownerID, format, date FROM photos INNER JOIN following ON ownerID = followedID WHERE followerID = ? ORDER BY date DESC LIMIT 50", requestingID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var stream []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.Owner, &photo.Format, &photo.Date)
		if err != nil {
			return nil, err
		}
		owner, _, _ := db.GetUserID(photo.Owner)
		photo.Username = owner.Username
		stream = append(stream, photo)
	}
	err = rows.Err()

	if rows.Err() != nil {
		return nil, err
	}

	return stream, err
}
