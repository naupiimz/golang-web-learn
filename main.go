package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Server is up and running")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Homepage"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
