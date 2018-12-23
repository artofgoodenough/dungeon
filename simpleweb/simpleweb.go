package main

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a very simple server in Golang"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	log.Println("Starting Server on :65432")
	err := http.ListenAndServe(":65432", mux)
	log.Fatal(err)
}
