package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/robfig/config"

	"log"
	"net/http"
)

type User struct {
	Username string
}

func LoginHandler(w http.ResponseWriter, r *http.Request, session sessions.Session, log *log.Logger) string {
	c, err := config.ReadDefault("users.cfg")

	if err != nil {
		return "Can't login. Problems reading user and password."
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	allowedUsername, _ := c.RawStringDefault("user")
	allowedPassword, _ := c.RawStringDefault("password")

	if username == allowedUsername && password == allowedPassword {

		log.Print("User WAS logged in.")

		session.Set("username", username)
		session.Set("password", password)
		http.Redirect(w, r, "/home", http.StatusFound)
		return "OK"
	}

	log.Print("User wasn't logged in. User " + username + " and password " + password)

	http.Redirect(w, r, "/login", http.StatusFound)
	return "Username or password incorrect"
}

func ShowWelcome(w http.ResponseWriter, user *User) string {
	w.Header().Set("Content-type", "text/html")
	return "Welcome back, " + user.Username + ". <a href='/logout'>Log out</a>"
}

func ProcessLogout(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
	session.Delete("username")
	http.Redirect(w, r, "/login", http.StatusFound)
	return "OK"
}

func Authorize(w http.ResponseWriter, r *http.Request, session sessions.Session, c martini.Context) {
	username := session.Get("username")
	if username == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	//if we found the user, let's create a new user struct and map it into the request context
	user := &User{}
	user.Username = username.(string)
	c.Map(user)
}
