package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/robfig/config"
)

type User struct {
	Username string
}

func main() {
	m := martini.Classic()

	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("auth"))
	m.Use(sessions.Sessions("go-webapp-part2", store))

	//on login, we will set the users username in the session
	m.Post("/login", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			username := r.FormValue("username")
			password := r.FormValue("password")
			session.Set("username", username)
			session.Set("password", password)
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})

	m.Post("/signup", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			username := r.FormValue("username")
			password := r.FormValue("password")
			c, err := config.ReadDefault("config.cfg")
			if err != nil {
				println("Error while loading config file!")
				http.Redirect(w, r, "/home", http.StatusFound)
				return "OK"
			}
			if c.HasSection(username) {
				println("Already registered!")
				http.Redirect(w, r, "/home", http.StatusFound)
				return "OK"
			}
			//add new username and password to config file
			c.AddSection(username)
			c.AddOption(username, "password", password)
			c.WriteFile("config.cfg", 0644, "")
			session.Set("username", username)
			session.Set("password", password)
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})

	//on logout, clear the username in the session
	m.Get("/logout", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			clearCookie(session)
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})

	//if the user has no username, we will redirect them to the login page
	m.Get("/home", authorize, func(user *User) string {
			return "Welcome back, " + user.Username
		})

	//redirect from root directory to "/home" page
	m.Get("/", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})

	m.Run()
}

func clearCookie(session sessions.Session) {
	session.Delete("username")
	session.Delete("password")
}

//The authorize middleware will search the session for a username
//if it doesnt find it, it will redirect to login
func authorize(w http.ResponseWriter, r *http.Request, session sessions.Session, c martini.Context) {
	username := session.Get("username")
	password := session.Get("password")
	//
	if username == nil || password == nil {
		clearCookie(session)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("Welcome, guest!<br/>You can <a href=\"signup\">sign up</a> or <a href=\"login\">log in</a> now!"))
		return
	}
	cf, err := config.ReadDefault("config.cfg")
	//problems with loading config files
	if err != nil {
		clearCookie(session)
		println("Error while loading config file!")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	cfgPassword, err := cf.String(username.(string), "password")
	//problems with reading password from config file
	if err != nil || cfgPassword != password.(string) {
		clearCookie(session)
		println("Bad login or password!")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	//if we found the user, let's create a new user struct and map it into the request context
	user := &User{}
	user.Username = username.(string)
	c.Map(user)
}
