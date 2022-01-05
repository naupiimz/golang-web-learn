package main

import (
	"goweb-learn/handler"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Server is up and running")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
