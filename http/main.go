package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var tpl *template.Template
var userData = make(map[string]User)

type PageData struct {
	EmailInvalid string
	PassInvalid  string
}

type User struct {
	Name     string
	Email    string
	Password string
}

var RandumNumber = rand.Intn(200)

func main() {
	var err error
	tpl, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		return
	}

	http.HandleFunc("/signup", handlefuncSignup)
	http.HandleFunc("/signuppost", signupmethod)
	http.HandleFunc("/login", loginfunc)
	http.HandleFunc("/loginpost", loginmethod)
	http.HandleFunc("/home", homefunc)
	http.HandleFunc("/logout", logoutfunc)
	http.HandleFunc("/", indexHandle)

	http.ListenAndServe(":8080", nil)
}

// Function to get the login cookie
func getLoginCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie("logincookie")
}

// Handle home page access and redirection based on login status
func homefunc(w http.ResponseWriter, r *http.Request) {
	cookie, err := getLoginCookie(r)
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// Default route that checks for login status and redirects accordingly
func indexHandle(w http.ResponseWriter, r *http.Request) {
	cookie, err := getLoginCookie(r)
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// Display the signup page
func handlefuncSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered the singup ")
	// Check login status to allow access
	cookie, err := getLoginCookie(r)
	if err == nil && cookie.Value != "" {
		// Optionally redirect or display message for logged-in users
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	err = tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		http.Error(w, "Signup page is not found: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handle signup form submission
func signupmethod(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Input validation
	if email == "" || password == "" || name == "" {
		tpl.ExecuteTemplate(w, "signup.html", "All fields are required")
		return
	}

	// Check if user already exists
	if _, ok := userData[email]; ok {
		tpl.ExecuteTemplate(w, "signup.html", "User already exists")
		return
	}

	// Register new user
	userData[email] = User{Name: name, Email: email, Password: password}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Handle login page request
func loginfunc(w http.ResponseWriter, r *http.Request) {
	cookie, err := getLoginCookie(r)
	if err == nil && cookie.Value == strconv.Itoa(RandumNumber) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

// Handle login form submission
func loginmethod(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in
	cookie, err := getLoginCookie(r)
	if err == nil && cookie.Value != "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailLogin")
	password := r.FormValue("passwordLogin")

	// Validate login form fields
	if email == "" {
		n := PageData{EmailInvalid: "Email is required"}
		tpl.ExecuteTemplate(w, "login.html", n)
		return
	} else if password == "" {
		n := PageData{PassInvalid: "Password is required"}
		tpl.ExecuteTemplate(w, "login.html", n)
		return
	}

	// Check if user exists
	user, ok := userData[email]
	if !ok || user.Password != password {
		n := PageData{PassInvalid: "Invalid credentials"}
		tpl.ExecuteTemplate(w, "login.html", n)
		return
	}

	// Create login cookie
	CookirForLogin := &http.Cookie{
		Name:   "logincookie",
		Value:  strconv.Itoa(RandumNumber),
		MaxAge: 300, // Cookie will expire in 5 minutes
		Path:   "/",
	}
	http.SetCookie(w, CookirForLogin)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// Handle logout request
func logoutfunc(w http.ResponseWriter, r *http.Request) {
	cookielogout := http.Cookie{
		Name:   "logincookie",
		Value:  "",
		MaxAge: -1, // Expire the cookie
		Path:   "/",
	}
	http.SetCookie(w, &cookielogout)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
