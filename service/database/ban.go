package database

func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err := db.c.Exec("INSERT INTO banned (bannerID,bannedID) VALUES (?, ?)", banner.ID, banned.ID)
	if err != nil {
		return err
	}
	return nil
}

// Database fuction che rimuovere un utente(banned)dalla lista dei banned di un'altro utente(banner)
func (db *appdbimpl) UnbanUser(banner User, banned User) error {

	_, err := db.c.Exec("DELETE FROM banned WHERE (bannerID = ? AND bannedID = ?)", banner.ID, banned.ID)
	if err != nil {
		return err
	}

	return nil
}

// [Util] Data base function per controllare se un utente è stato bannato.
// Restituisco 'true' se è banned, sennò 'false'
func (db *appdbimpl) BannedUserCheck(banner User, banned User) (bool, error) {
	// Utilizza il metodo QueryRow per eseguire una query SQL SELECT COUNT(*) che
	// conta quante volte l'utente banner appare nella tabella banned_users come utente bannato da banned
	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM banned WHERE bannedID = ? AND bannerID = ?",
		banned.ID, banner.ID).Scan(&cnt)

	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true, err
	}

	// If the counter is 1 then the user was banned
	if cnt == 1 {
		return true, nil
	}
	return false, nil
}

func (db *appdbimpl) NoMoreFriends(banner User, banned User) (err error) {
	query1 := "DELETE FROM following WHERE followerID = ? AND followedID = ?;"
	_, err = db.c.Exec(query1, banner.ID, banned.ID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query1, banned.ID, banner.ID)

	if err != nil {
		return err
	}

	query2 := "DELETE FROM likes WHERE likerID = ? AND photoID IN (SELECT photoID FROM photos WHERE ownerID = ?);"
	_, err = db.c.Exec(query2, banned.ID, banner.ID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query2, banner.ID, banned.ID)
	if err != nil {
		return err
	}
	return nil
}
