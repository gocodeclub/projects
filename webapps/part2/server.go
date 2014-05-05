package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

type User struct {
	Username string
}

func main() {
	m := martini.Classic()

	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("secret-pass"))
	m.Use(sessions.Sessions("go-webapp", store))

	//on login, we will set the users username in the session
	m.Post("/login", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
		username := r.FormValue("username")
		session.Set("username", username)
		http.Redirect(w, r, "/home", http.StatusFound)
		return "OK"
	})

	//if the user has no username, we will redirect them to the login page
	m.Get("/home", authorize, func(user *User) string {
		return "Welcome back, " + user.Username
	})

	//on logout, clear the username in the session
	m.Get("/logout", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
		session.Delete("username")
		http.Redirect(w, r, "/login", http.StatusFound)
		return "OK"
	})

	m.Run()
}

//The authorize middleware will search the session for a username
//if it doesnt find it, it will redirect to login
func authorize(w http.ResponseWriter, r *http.Request, session sessions.Session, c martini.Context) {
	username := session.Get("username")
	if username == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	//if we found the user, let's create a new user struct and map it into the request context
	user := &User{}
	user.Username = username.(string)
	c.Map(user)
}
