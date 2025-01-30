package handler

/*func ErrorHandler(w http.ResponseWriter, status int, errMsg string, description string) {
	// Parse
	t, err := template.ParseFiles("../client/templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Fill struct info
	FullError := struct {
		StatusCode   string
		ErrorMessage string
		Description  string
	}{
		StatusCode:   fmt.Sprintf("%v", status),
		ErrorMessage: errMsg,
		Description:  description,
	}

	// write header status code and execute page
	w.WriteHeader(status)
	err = t.Execute(w, FullError)
	if err != nil {
		fmt.Println("Error Executing Error Page", err)
		return
	}
}*/
