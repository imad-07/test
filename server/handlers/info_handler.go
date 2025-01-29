package handler

import (
	"database/sql"
	"net/http"
	"time"

	"forum/server/data"
	"forum/server/helpers"
	"forum/server/service"
	"forum/server/shareddata"

	"github.com/mattn/go-sqlite3"
)

type InfoHandler struct {
	InfoService service.InfoService
}

func NewInfoHandler(db *sql.DB) *InfoHandler {
	infoData := data.InfoData{
		Db: db,
	}

	infoService := service.InfoService{
		InfoData: infoData,
	}

	return &InfoHandler{
		InfoService: infoService,
	}
}

func (h *InfoHandler) Info(w http.ResponseWriter, r *http.Request) {
	// parse user uid
	uid := ""
	userUID, errCookie := r.Cookie(shareddata.SessionName)
	if errCookie == nil {
		uid = userUID.Value
	}

	// Get Info Data
	infoData, err := h.InfoService.GetInfoData(uid)
	if err != nil {
		if err == sqlite3.ErrLocked {
			helpers.WriteJson(w, http.StatusLocked, struct {
				Message string `json:"message"`
			}{Message: "Database Locked"})
				return
		}

		helpers.WriteJson(w, http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{Message: "Internal Server Error"})
		return
	}
	if errCookie == nil && !helpers.CheckExpiredCookie(userUID.Value, time.Now(), h.InfoService.InfoData.Db) {
		infoData.Authorize = false
		DeleteSessionCookie(w, userUID.Value)
	}

	helpers.WriteJson(w, http.StatusOK, infoData)
}
