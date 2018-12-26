package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	str := `<html>
				<head>
					<title>A Simple Web Server Example</title>
				</head>
				<body>
					<h1>Hello from a very simple server in Golang</h1>
				</body>
			</html>`
	w.Write([]byte(str))
}

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
}

func writeError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Some Error Message here ....")
}

func redirectHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type SampleJson struct {
	Name    string
	Age     int8
	Hobbies []string
	Height  float32
}

func writeJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sample := SampleJson{
		Name:    "John Wayne",
		Age:     102,
		Hobbies: []string{"swimming", "tennis", "bowling", "hiking"},
		Height:  6.75,
	}
	jsonText, _ := json.Marshal(&sample)
	w.Write(jsonText)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/error", writeError)
	mux.HandleFunc("/redirect", redirectHeader)
	mux.HandleFunc("/json", writeJson)
	log.Println("Starting Server on :65432")
	err := http.ListenAndServe(":65432", mux)
	log.Fatal(err)
}
