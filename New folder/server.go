package main

import (
	"html/template"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Serve the home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	})

	// Serve the /res page
	http.HandleFunc("/res", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("reservation.html"))
		tmpl.Execute(w, nil)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
