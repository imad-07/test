package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"forum/server/data"
	"forum/server/helpers"
	postapp "forum/server/service"
	"forum/server/shareddata"
)

type ReactHandler struct {
	ReactService postapp.ReactService
}

func NewReactHandler(db *sql.DB) *ReactHandler {
	data := data.ReactionDB{
		DB: db,
	}

	reactService := postapp.ReactService{
		ReactData: data,
	}

	return &ReactHandler{ReactService: reactService}
}

func (service *ReactHandler) ReactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	react := shareddata.React{}
	json.NewDecoder(r.Body).Decode(&react)
	cookie, err := r.Cookie(shareddata.SessionName)
	if err != nil || !helpers.CheckExpiredCookie(cookie.Value, time.Now(), service.ReactService.ReactData.DB) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, id := postapp.GetUser(service.ReactService.ReactData.DB, cookie.Value)

	err = service.ReactService.CheckReactInput(react)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = service.ReactService.ReactionService(react, id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, err := service.ReactService.LikesTotal(react.Thread_type, react.Thread_id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.IsLiked, response.IsDisliked = service.ReactService.GetLikedThread(react.Thread_type, react.Thread_id, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
