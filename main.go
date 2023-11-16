package main

import (
	"fmt"
	"net/http"
	"os"

	endpoints "groupieTrecker/internal/endpoints"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	styles := http.FileServer(http.Dir("internal/templates/css"))

	router := http.NewServeMux()
	router.Handle("/css/", http.StripPrefix("/css/", styles))
	router.HandleFunc("/", endpoints.GetMainPage)
	router.HandleFunc("/artist", endpoints.GetArtist)
	router.HandleFunc("/search", endpoints.SearchHandler)
	fmt.Println("listening on: http://localhost:" + port + "/")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error: several connecrtions")
		return
	}
}
