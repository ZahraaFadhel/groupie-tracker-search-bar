package main

import (
	"fmt"
	groupieTracker "groupieTracker/handlers"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", getHandler)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/style.css")
	})

	port := "localhost:8080"

	fmt.Printf("Server is working on http://" + port + "\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		groupieTracker.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		groupieTracker.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		groupieTracker.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
