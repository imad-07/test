package handler

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		//ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Maybe GET Method Will Work!")
		return
	}
	ServeHomePage(w, r)
}

// Helper
func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//ErrorHandler(w, http.StatusNotFound, "Page Not Found", "Page You Are Looking For Doesn't Exist")
		return
	}
	t, err := template.ParseFiles("../ui/templates/index.html")
	if err != nil {
		//ErrorHandler(w, http.StatusInternalServerError, "inernal Server Error", "Error While Parsing index.html")
		log.Println("Unexpected error", err)
		return
	}
	t.Execute(w, nil)
}
