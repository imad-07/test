package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("index.html", "page2.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", asciiArtHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")

	fmt.Println(text)
	banner := r.FormValue("banner")
	fmt.Println(banner)

	// Here you would call the ASCII art generation function
	// For now, let's assume we have a function `generateASCIIArt(text, banner string) string`
	result := generateASCIIArt(text, banner)

	data := struct {
		Text   string
		Banner string
		Result string
	}{
		Text:   text,
		Banner: banner,
		Result: result,
	}
	err := templates.ExecuteTemplate(w, "page2.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateASCIIArt(text, banner string) string {
	return text + " " + banner
}
