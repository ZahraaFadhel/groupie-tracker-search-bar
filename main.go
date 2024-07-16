package main

import (
	"fmt"
	groupieTracker "groupieTracker/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", groupieTracker.GetHandler)
	http.HandleFunc("/artist/", groupieTracker.GetArtist)
	http.HandleFunc("/aboutUs/", groupieTracker.AboutUsHandler)

	http.HandleFunc("/styleArtists.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/artists/styleArtists.css")
	
	})
	http.HandleFunc("/styleArtist.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/artist/styleArtist.css")
	})
	http.HandleFunc("/error/error.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "error/error.css")
	})
	http.HandleFunc("/aboutUs.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/aboutUs/aboutUs.css")
	})
	http.HandleFunc("/search.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/artists/search.js")
		// enableCors(&w)
	})

	port := "localhost:8080"
	fmt.Printf("Server is working on http://" + port + "\n")
	
	err := http.ListenAndServe(":8080", nil)
	if err == nil {
		groupieTracker.OpenBrowser("http://localhost:8080")
	}
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }