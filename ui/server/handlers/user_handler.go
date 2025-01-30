package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"forum/server/data"
	"forum/server/service"
	"forum/server/shareddata"

	"github.com/mattn/go-sqlite3"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	userData := data.UserData{
		DB: db,
	}

	userService := service.UserService{
		UserData: &userData,
	}

	return &UserHandler{
		UserService: &userService,
	}
}

func (h *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//	ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Maybe GET Method Will Work!")
		return
	}

	// Parse Data
	var user shareddata.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Proccess Data and Insert it
	err = h.UserService.RegisterUser(&user)
	if err != nil {
		if err == sqlite3.ErrLocked {
			http.Error(w, "Database Is Busy!", http.StatusLocked)
			return
		}
		// Username
		if err.Error() == shareddata.Errors.InvalidUsername {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Password
		if err.Error() == shareddata.Errors.InvalidPassword {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Email
		if err.Error() == shareddata.Errors.InvalidEmail {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err.Error() == shareddata.Errors.LongEmail {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// General
		if err.Error() == shareddata.Errors.UserAlreadyExist {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Error While Registering The User.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("You Logged In Successfuly!"))
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}

/*func (h *UserHandler) ServeSignUpPage(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Maybe GET Method Will Work!")
		return
	}

	// Check if user already exist (if user id != 0 means user exist)
	uid, err := r.Cookie(shareddata.SessionName)
	if err == nil {
		_, id := service.GetUser(h.UserService.UserData.DB, uid.Value)
		if id != 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	// Parse
	t, err := template.ParseFiles("../client/templates/signup.html")
	if err != nil {
		//ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Error While Parsing signup.html.")
		return
	}

	// Execute
	err = t.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Error While Executing signup.html")
		return
	}
}*/

// Login
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Maybe POST Method Will Work!")
		return
	}

	// Parse Data
	var user shareddata.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Proccess Data and Insert it
	err = h.UserService.LoginUser(&user)
	if err != nil {
		if err == sqlite3.ErrLocked {
			http.Error(w, "Database Is Busy!", http.StatusLocked)
			return
		}
		// Email
		if err.Error() == shareddata.Errors.InvalidEmail {
			fmt.Println(2)
			http.Error(w, shareddata.Errors.InvalidEmail, http.StatusBadRequest)
			return
		}
		if err.Error() == shareddata.Errors.LongEmail {
			fmt.Println(3)
			http.Error(w, shareddata.Errors.LongEmail, http.StatusBadRequest)
			return
		}

		// Password
		if err.Error() == shareddata.Errors.InvalidPassword {
			fmt.Println(4)
			http.Error(w, shareddata.Errors.InvalidPassword, http.StatusBadRequest)
			return
		}
		// General: User Doesn't Exist
		if err.Error() == shareddata.Errors.InvalidCredentials {
			http.Error(w, shareddata.Errors.InvalidCredentials, http.StatusUnauthorized)
			return
		}

		if err == sql.ErrNoRows {
			http.Error(w, shareddata.Errors.InvalidCredentials, http.StatusUnauthorized)
			return
		}

		log.Println("Unexpected error:", err)
		http.Error(w, "Error While logging To An  Account.", http.StatusInternalServerError)
		return
	}
	SetSessionCookie(w, user.Uuid)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You Logged In Successfuly!"))
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*func (h *UserHandler) ServeLoginPage(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		//ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Maybe GET Method Will Work!")
		return
	}

	// Check if user already exist (if user id != 0 means user exist)
	uid, err := r.Cookie(shareddata.SessionName)
	if err == nil {
		_, id := service.GetUser(h.UserService.UserData.DB, uid.Value)
		if id != 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	// Parse
	t, err := template.ParseFiles("../ui/templates/index.html")
	if err != nil {
		//ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Error While Parsing login.html")
		return
	}

	// Execute
	err = t.Execute(w, nil)
	if err != nil {
		//ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Error While Executing login.html")
		return
	}
}*/

// Helpers
func SetSessionCookie(w http.ResponseWriter, uid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   shareddata.SessionName,
		Value:  uid,
		Path:   "/",
		MaxAge: 3600,
	})
}

func DeleteSessionCookie(w http.ResponseWriter, uid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   shareddata.SessionName,
		Value:  uid,
		Path:   "/",
		MaxAge: -1,
	})
}
