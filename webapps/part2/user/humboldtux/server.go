package main

import (
	"log"
	"net/http"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/spf13/viper"
)

type User struct {
	Username string
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	//Load configuration
	initializeConfig()

	m := martini.Classic()

	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("secret-pass"))
	m.Use(sessions.Sessions("go-webapp", store))
	// map template rendering into martini, with layout template
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	//if the user has no username, we will redirect them to the login page
	m.Get("/home", authorize, Home)
	m.Get("/login", Login)
	//on login, we will set the users username in the session
	m.Post("/login", binding.Bind(LoginForm{}), PostLogin)
	//on logout, clear the username in the session
	m.Get("/logout", Logout)

	m.Run()
}

func initializeConfig() {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatal("Problem loading configuration file")
	}
	viper.SetDefault("admin", "default")
}

func Home(ren render.Render, user *User) {
	ren.HTML(200, "home", user)
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

func Login(ren render.Render) {
	ren.HTML(200, "login", nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request, session sessions.Session, login LoginForm) string {
	hash := viper.Get(login.Username).(string)

	//if hash != login.Password {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(login.Password)) != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return "KO"
	} else {
		session.Set("username", login.Username)
		http.Redirect(w, r, "/home", http.StatusFound)
		return "OK"
	}
}

//on logout, clear the username in the session
func Logout(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
	session.Delete("username")
	http.Redirect(w, r, "/login", http.StatusFound)
	return "OK"
}
