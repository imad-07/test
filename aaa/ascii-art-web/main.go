package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("index.html", "result.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if false {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Bad Request
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "Missing something", http.StatusNotFound)
		return
	}
	result := generate(text, banner)
	data := struct {
		Result string
	}{
		Result: result,
	}
		err := templates.ExecuteTemplate(w, "result.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
	}
}

func generate(str string, banner string) string {
	input := ""
	for i := 0; i < len(str); i++ {
		if str[i] != 13 {
			input = input + string(str[i])
		}
	}
	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			continue
		}
		if input[i] < 32 || input[i] > 128 {
			return "wrong input"
		}
	}
	word := split(input)
	BANNER := banner
	if BANNER == "" {
		BANNER = "standard.txt"
	}
	fileContent, err := os.ReadFile(BANNER)
	if err != nil {
		return "erro in file\n"
	}
	lettres := getLettres(fileContent)
	return output(lettres, word)

}
func output(lettres [][]string, word []string) string {
	output := ""
	for l := 0; l < len(word); l++ {
		if word[l] == "\n" {
			output += "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				output += lettres[word[l][j]-32][i]
			}
			output += "\n"
		}
	}
	return output
}

func split(str string) []string {
	word := ""
	splitedword := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	if word != "" {
		splitedword = append(splitedword, word)
	}
	return splitedword
}
func getLettres(fileContent []byte) [][]string {
	Content := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			Content = append(Content, fileContent[i])
		}
	}
	fileContent = Content
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if i != len(fileContent)-1 && fileContent[i] == '\n' && fileContent[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if fileContent[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, fileContent[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}
