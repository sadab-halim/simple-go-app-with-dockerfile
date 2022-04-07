package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go application running inside Docker.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Okay\n")
}

func main() {
	fmt.Println("Web Server is starting on port 8080.")
	http.HandleFunc("/". indexHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}