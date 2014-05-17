package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetHandler(operand func(int, int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received: ", r.URL.String())
		values := r.URL.Query()
		if a, err := strconv.Atoi(values.Get("a")); err == nil {
			if b, err := strconv.Atoi(values.Get("b")); err == nil {
				result := fmt.Sprint("result: ", operand(a, b))
				w.Write([]byte(result))
				return
			}
		}
		w.Write([]byte("Failed to parse one of the args"))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Code Club"))
}

func main() {
	log.Println("listening...")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/add", GetHandler(func(a int, b int) int {
		return a + b
	}))
	http.HandleFunc("/sub", GetHandler(func(a int, b int) int {
		return a - b
	}))
	http.HandleFunc("/mul", GetHandler(func(a int, b int) int {
		return a * b
	}))
	http.HandleFunc("/div", GetHandler(func(a int, b int) int {
		return a / b
	}))
	http.ListenAndServe(":3000", nil)
}
