package main

import (
	"./handlers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()

	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("secret-pass"))
	m.Use(sessions.Sessions("go-webapp", store))

	//on login, we will set the users username in the session
	m.Post("/login", handlers.LoginHandler)

	//if the user has no username, we will redirect them to the login page
	m.Get("/home", handlers.Authorize, handlers.ShowWelcome)

	//on logout, clear the username in the session
	m.Get("/logout", handlers.ProcessLogout)

	m.Run()
}
