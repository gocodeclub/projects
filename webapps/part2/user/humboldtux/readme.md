#Webapps in Go: Part 2

###Summary

- What's Dependency Injection?
- A closer look at Martini
- Creating your first custom middleware

Welcome back to the second edition of the webapp series here at the Go Code Club.  Hopefully all of you have completed your first assignment, and are thirsty for the next one!  This week's assignment should prove a bit more difficult, so get ready!  
This week, we are continuing our conversation about webapps in Go. Specifically, we are going to pay some attention to one of my favorite Golang projects, Martini.  

###So, whats in a Martini?

[Jeremy Saenz](https://github.com/codegangsta), the creator of Martini said that he [named it Martini](http://thechangelog.com/117/) because these days, there are so many ways to make a martini.  Similarly, there are many ways to make a webapp as well.  Everyone throws different things into them.

What makes Martini special?  We learned last week that it's easy enough to handle HTTP Requests in Go, so why do we need Martini?  Well first of all, Martini gives us a few nice-to-haves.  We can handle HTTP requests, based on the the HTTP verb, by simply passing a handler function.  This concept is called *routing*, and is probably the most common feature of web frameworks.  It also comes with a set of extra's that you can apply, optionally of course, in the [Martini-Contrib repo](https://github.com/martini-contrib).

###The Killer Feature

In my opinion, the real killer feature of Martini however, is its method of dependency injection.  Dependency Injection, for those of you who don't know, is a concept that was created to help manage complex systems.  Heres a great definition of Dependency Injection from [Stack Overflow](http://stackoverflow.com/questions/130794/what-is-dependency-injection):

>"Dependency Injection" is a 25-dollar term for a 5-cent concept. [...] Dependency injection means giving an object its instance variables. [...].

Before Martini, I had always kind of shied away from the word *dependency injection*.  It honestly sounds scary, but as James Shore points out above, it's kind of an over hyped word.  For now, lets just think of it as a technique used to simplify working with multiple components.  You'll see in a second that its not as crazy as it sounds.

###Martini Core

At the core of understanding everything that's going on in Martini, is two concepts:  **Handlers** and **Services**.  

**Handlers** are a pretty basic component of web frameworks, as mentioned above.  In Martini, a handler is basically any *callable function*.  A basic example of how you might use a handler is this.  

Let's say you want to display the home page on your website.  Basically, when the user is directed to `/home`, we want to take a certain *action*.  This *action* is essentially our handler, and `/home` is our route.

```javascript
m.Get("/home", func() string {
  return "Welcome Home!" // HTTP 200 : "Welcome Home!"
})
```

Pretty simple, right?  Okay, lets take things one step further.  Handlers can actually be *stacked*.  That means, we can provide multiple functions that will be run, called in the order we pass them.  Why would we want to have multiple handlers, you might ask?  For those of you who are relatively new to writing web apps, this concept is often referred to as the middleware layer.  

**Middleware** is another fancy term that describes something fairly simple.  In a web app, you often want to do certain things before, or after, the request.  For example, you may want to authorize the user before you display the page, or redirect them to the login screen instead.

One thing to note is that only one handler can write a response (actually physically return something to the page).  Usually, that's your last handler, but not always.  Let's take another contrived example:

```javascript
m.Get("/home", authorize, func() {
  // this will execute as long as authorize doesn't write a response
  return "Welcome Home!" // HTTP 200 : "Welcome Home!"
}) 
```

This is almost identical to the first example, except this example takes one additional handler.  In this case, we have a *middleware* that insures the person loading the page is authenticated. 

####Side Note
...But wait, how many parameters does this martini.Get() function take?  First we passed in one handler, now we're passing in two.  What gives?  This is a cool feature about Go, called [Variadic Functions](http://www.golang-book.com/7#section3).  If you pop open router.go in the martini package, you'll see the Router interface:

```javascript
type Router interface {
	...
	// Get adds a route for a HTTP GET request to the specified matching pattern.
	Get(string, ...Handler) Route
	...
}
```

Cool, the "..." before the type specifies that we will pass in *0 or more arguments*. For Martini though, we need to pass in at least 1 or else it will panic.  


###Services

You now have a pretty good grasp on Handlers:  A Handler is basically just a function that does something.  Well, like all good functions, they're allowed to take parameters.  A parameters that a Martini Handler takes is called a **Service**.

Once again, we have another hyped up word.  Services are just a plain old parameter to a function.  However, there is a bit going on behind the scenes that make this all possible.  This is where the dependency injection comes in.

In short, you can map parameters to your martini instance, making them optionally available as parameters in any and every handler function.  I think of it kind of like a grab bag of params you might need.  You can put things into the grab bag by calling `Map()`, and get them out of the bag by including them as a parameter to your function later.

By default, a classic Martini has a few services mapped for you, like http.ResponseWriter, http.Request and martini.Context, so you can feel free to include them as params when needed.  It's easy, however, to map your own services.  Let's write the authorize function we referenced above:

```javascript
type User struct {
	Username string
}
...

//The authorize middleware will search the session for a username
//if it doesn't find it, it will redirect to login
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
```

In this example, we're going to pull in the sessions package from Martini contrib.  This is a vastly oversimplified approach, but we will improve it next week.  For now, we are going to simply check to see if there is a value in the username session.  If we find it, we're going to create a new instance of a user, and assign the username.  We then map the user object, which is going to make it available for us later.  

```javascript
m.Get("/home", authorize, func(user *User) string {
	return "Welcome back, " + user.Username
})
```

Services can be mapped globally on the Martini instance, or mapped to the martini.Context object to map it at the request level.  You would want to map something globally if you wanted it in every request, for example a database connection object.  In our case, we map the user to the request, because the object will depend on which user is loading the page.

I've added a `*User` param to our handler for `/home`.  I can safely do this because I know that authorize will either return a user, or redirect, which would stop us from getting that far.  *This is dependency injection at its best*.  The concept is that our handler function doesn't need to know how or where the user came from.  When asked for, martini will *resolve* the param using reflection, and return you the value.  This is a simple example, but as your application grows, this technique can vastly simplify your code.

I've added two more routes to handle adding the session and deleting it, so you can test it out in your browser.  

```javascript
//on login, we will set the users username in the session
m.Post("/login", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
	username := r.FormValue("username")
	session.Set("username", username)
	http.Redirect(w, r, "/home", http.StatusFound)
	return "OK"
})
...
//on logout, clear the username in the session
m.Get("/logout", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
	session.Delete("username")
	http.Redirect(w, r, "/login", http.StatusFound)
	return "OK"
})
```

###Your assignment

This week, we learned how to write your own handlers, and how to map services.  This week, I'm going to have you play around with Martini a bit more.

**Assignment #2**

- Run the example (full source on github), and test it out in your browser.  Try going to /home, and notice how you get redirected.  Then, login, and see what happens.  To test it out again, you can visit /logout.

- Then, extend the example above to actually do something.  What it does is up to you, this is where you can get creative.  If you can't think of anything terribly exciting, start by looking at the [Martini Contib repo](https://github.com/martini-contrib).  Render and/or Binding are a good place to start.

- Send a PR with your implementation.  Take a look [here](https://github.com/gocodeclub/projects/blob/master/README.md) for instructions on how to format your contribution.

- Leave a comment here, if you'd like, explaining what you've written, and a link to your folder in the repo.

As always, if I missed something or made a mistake, feel free to send me a PR
https://github.com/gocodeclub/projects/tree/master/webapps/part2

