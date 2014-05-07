package main

import(
	"net/http"
	"log"
	"fmt"
)


func main() {

	http.HandleFunc("/hello", hello)
	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
	
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	
	if name == "" {
		name = "Go Code Club"
	}

	formHTML := "<form>Enter your name: <input name=\"name\" type=\"text\"/><input type=\"submit\" value=\"Submit\"/></form><br/>"
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s Hello %s!", formHTML, name)
}

