package helpers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Function help in json writing
// write a json response
func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Unexpected Error %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func CheckExpiredCookie(uid string, date time.Time, db *sql.DB) bool {
	var expired time.Time
	db.QueryRow("SELECT expired_at FROM user_profile WHERE uid = ?", uid).Scan(&expired)

	return date.Compare(expired) <= -1
}
