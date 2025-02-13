package main

import (
	"fmt"
	"net/http"

	"Student/functions"
)

func main() {
	http.HandleFunc("/", functions.MainPage)
	http.HandleFunc("/static/", functions.StyleHandler)
	http.HandleFunc("/ascii-art", functions.GenerateHandler)
	fmt.Println("Server is running at http://localhost:8020")
	http.ListenAndServe(":8020", nil)
}
