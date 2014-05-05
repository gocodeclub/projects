#Golang webapps, pt 1

###Summary

- Native Go HTTP
- Web Frameworks in Go
- Introduction to Martini

**Note** *If you are completely new to go, start with [this](http://app.gocodeclub.com/t/gettings-started-with-go/12) tutorial on how to get go setup, and move on to this assignment*


###Webapps in Go: Part 1

Welcome to the first edition of the first series at the Go Code Club!  The first topic we are going to be covering here is an area that I am fortunate to have some good experience with, so I found it a good starting spot. And really, how could you have a club about Go and NOT talk about webapps ;)

Ok, let's say you want to create a webapp that *foos bars*.  Where do you start? Which framework do you choose?  There are some great options out there, and we will cover several in this series.  However, I want to stress that before you start slinging code with your favorite framework, you should fully understand how the Net/HTTP libraries work.  You might be happily surprised by how much heavy lifting it does all by itself.

###Hello net/http

Let's start with a simple example, using only net/http.  Out of the box, Go can do some pretty complicated things, like [handle http](http://golang.org/pkg/net/http/), [serve files](http://golang.org/pkg/net/http/#FileServer), [parse request bodies](http://golang.org/pkg/encoding/json/), and much more.  

Here we have an example that listens on port 3000 and waits for requests.  

```go
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
```

Pretty basic stuff, but I wanted to point a few things out.  The `http.HandleFunc` method is a bit of shorthand that actually does a few things.  It takes a function that conforms to the [Handler](http://golang.org/pkg/net/http/#Handler) interface, which is basically a function that takes `(ResponseWriter, *Request)` as parameters.  

It also uses the DefaultServerMux to multiplex the connection.  If you want to know more about how the Handler interface works, check out this [really detailed article](http://www.alexedwards.net/blog/a-recap-of-request-handling) on the subject.

###Web Frameworks in Go

There are a **ton** of great web frameworks right now for Go, to name a few: [Revel](http://revel.github.io/), [Martini](http://martini.codegangsta.io/), [Beego](http://beego.me/), [Gorilla](http://www.gorillatoolkit.org/), and many more.  I hope by the end of this series, you'll have a better idea of how to choose what framework is for you, and to understand the basic value propositions each give. 

I think first, you need to ask yourself, what's the point of a web framework, or a framework in general? 
   - **frameworks make you productive**.  Frameworks ideally keep you writing application logic, and not worrying about the semantics of a language.  Some larger frameworks even give you a way of organizing your code, and other niceties.
   - **frameworks solve common problems** - for example, you want to provide CSFR support?  Or maybe you want to handle basic auth?  These are *very* common features of a web framework. 

This all sounds great, but there are a few reasons you might want to be careful in choosing a framework.  
   - **Frameworks abstract details from you**:  This is both good and bad.  Today, you might not care about some low level language details, but tomorrow you might have a new feature that requires greater control.  
   - **Your needs are unique** - at some point, you will run into something that the framework doesn't handle very gracefully.  You're then going to have to ask yourself, "how to I solve X issue in Y framework" rather than "how do I solve X issue in Y language".

If you decide that you need some additional features, and that a framework could help you, you may want to check one of these out.  If not, Net/HTTP can take you a long way.

Some frameworks are very minimalistic, like Gorilla, and Martini, while others are more full feature, like Revel.  We will cover more of the pros and cons to each as we go, but first, I want to dive into a very popular framework right now called Martini.

###Martini, anyone?

```go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```

Martini is a minimalistic web framework for quickly writing modular web applications/services in Go.  Whats unique about Martini is its method of dependency injection to deal with common request issues.  

To get this hello world up and running, you need to first copy the code and place it in your GOPATH, lets call it server.go.  

For martini to work, we need to get its dependencies.  Run the following command to get the martini source:

`go get github.com/go-martini/martini`

And then from the directory that you put the `server.go` file, run `go run server.go`.  Great!  You now have a server running on port 3000.  To test everything out, go to your browser and type in `localhost:3000`.  Press enter and you should see "Hello World".  

###Your assignment

This week, I'm going easy on you.  I want to cater to all levels, some of whom might be just starting with Go, so this assignment will vary depending on where you're at.

**Your Task**

Run both of the hello world examples above (the standard net/http one, and the martini one).  Try playing around with adding routes, parameters, or even POSTing some params. 

If your completely new to Go, read [my getting started post](http://app.gocodeclub.com/t/gettings-started-with-go/12), get Go installed and setup, and then do this assignment.

For those of you who have already played around with Martini, great!  You're one step ahead.  Make sure you're up-to-date with the latest Martini docs. Also make sure you're linking to the most recent version of Martini, as [its changed](https://groups.google.com/forum/#!topic/martini-go/hw8trVJXB2c). Next week's assignment will be fun for everyone at all levels.  

###Next week!
The next post will be covering a deeper dive into Martini to go through exactly how this all works.  Here's a sneak peek at what I have planned:

- What's Dependency Injection?
- What's Reflection, and how is it used?
- Creating your first custom middleware


PS: I am hosting a copy of each tutorial on Github, along with its relevant examples.  If there is a mistake, or you feel I could have elaborated in an area, feel free to send a pull request :) 
https://github.com/gocodeclub/projects


