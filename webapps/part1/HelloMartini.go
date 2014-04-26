package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "<h1>Hello World!</h1>"
	})

	m.Get("/hello:name", func(params martini.Params) string {
		return "<h1>Hello " + params["name"] + "</h1>"
	})

	m.Get("/hello/:**", func(params martini.Params) string {
		return "<h1>Hello " + params["_1"] + "</h1>"
	})

	m.Run()
}
