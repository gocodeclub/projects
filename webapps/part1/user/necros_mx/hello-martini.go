package main

import (
	"github.com/go-martini/martini"
)


func main() {
	m := martini.Classic()

	m.Get("/hello/:name", func (params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Get("/hello", func () string {
		return "Hello Go Code Club!"
	})

	m.Run()
}
