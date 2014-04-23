package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go Code Club!"))
}
