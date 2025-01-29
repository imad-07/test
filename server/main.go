package main

import (
	"fmt"
	"net/http"

	"forum/server/data"
	"forum/server/routes"
)

func main() {
	db, err := data.OpenDb()
	if err != nil {
		fmt.Println("Error in opening of database:", err)
		return
	}

	err = data.InitTables(db)
	if err != nil {
		fmt.Println("Error in intializing of tables:", err)
		return
	}

	server := http.Server{
		Addr:    ":8081",
		Handler: routes.Routes(db),
	}

	fmt.Println("http://localhost:8081/")

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error in starting of server:", err)
		return
	}
}
