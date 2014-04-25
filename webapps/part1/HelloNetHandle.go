package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //Analytical parameters，The default will not do
	fmt.Println(r.Form) //The output information to the server
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello World!") //This is written to the w output to the client
}

func main() {
	http.HandleFunc("/", sayHelloWorld)      //Set access route
	err := http.ListenAndServe(":8081", nil) //Set the listener port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
