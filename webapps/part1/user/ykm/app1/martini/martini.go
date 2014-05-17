package main

import (
	"github.com/go-martini/martini"
	"log"
)

func main() {
	log.Println("listening...")
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello world"
	})
	m.Get("/hello/:name", func(params martini.Params) string {
		return "hello " + params["name"]
	})
	m.Run()
}
