package database

func (db *appdbimpl) FollowUser(followerID string, followedID string) (err error) {

	query := "INSERT INTO following (followerID, followedID) VALUES (?,?);"

	_, err = db.c.Exec(query, followerID, followedID)
	return err
}

func (db *appdbimpl) UnfollowUser(followerID string, followedID string) (err error) {

	query := "DELETE FROM following WHERE (followerID = ? AND followedID = ?);"

	_, err = db.c.Exec(query, followerID, followedID)
	return err
}

func (db *appdbimpl) CheckFollow(followerID string, followedID string) (isFollowing bool, err error) {

	query := "SELECT EXISTS (SELECT '_' FROM following WHERE followerID = ? AND followedID = ?);"

	err = db.c.QueryRow(query, followerID, followedID).Scan(&isFollowing)
	return
}

func (db *appdbimpl) RemoveFollowBothDirections(user1ID string, user2ID string) (err error) {
	err = db.UnfollowUser(user1ID, user2ID)
	if err != nil {
		return err
	}
	err = db.UnfollowUser(user2ID, user1ID)
	return err
}
