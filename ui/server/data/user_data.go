package data

import (
	"database/sql"
	"time"

	"forum/server/shareddata"
)

type UserData struct {
	DB *sql.DB
}

func (database *UserData) InsertUser(user shareddata.User) error {
	_, err := database.DB.Exec("INSERT INTO user_profile (username, email, password, uid) VALUES (?, ?, ?, ?)",
		user.Username, user.Email, user.Password, user.Uuid)
	return err
}

func (database *UserData) CheckIfUserExists(username, email string) bool {
	var uname string
	var uemail string
	database.DB.QueryRow("SELECT username, email FROM user_profile WHERE username = ? OR email = ?",
		username, email).Scan(&uname, &uemail)
	return uname == username || uemail == email
}

func (database *UserData) GetUserPassword(email string) (string, error) {
	var password string
	err := database.DB.QueryRow("SELECT password FROM user_profile WHERE email = ?",
		email).Scan(&password)
	return password, err
}

func (database *UserData) UpdateUuid(uuid, email string) error {
	expire := time.Now().Add(time.Hour)
	_, err := database.DB.Exec("UPDATE user_profile SET uid = ?, expired_at = ? WHERE email = ?", uuid, expire, email)
	return err
}

func (database *UserData) GetUserName_Id(uuid string) (string, int) {
	var id int
	var username string
	err := database.DB.QueryRow("SELECT id, username FROM user_profile WHERE uid = ?",
		uuid).Scan(&id, &username)
	if err != nil {
		return "", 0
	}
	return username, id
}

func (database *UserData) CheckPassword(email, password string) bool {
	err := database.DB.QueryRow("SELECT email FROM user_profile WHERE email = ? AND password = ?",
		email, password).Scan(&email)
	return err == nil
}



// Helpers
