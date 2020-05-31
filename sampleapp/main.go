package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong Pong!")
	})

	http.HandleFunc("/showconfig", func(w http.ResponseWriter, r *http.Request) {
		value := os.Getenv("MESSAGE")
		fmt.Fprintf(w, value)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
