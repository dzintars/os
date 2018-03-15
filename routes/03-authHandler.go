package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func authHandler(r *mux.Router) {
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
	r.HandleFunc("/signup", signupPostHandler).Methods("POST")
	r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	r.HandleFunc("/signin", signinPostHandler).Methods("POST")
	r.HandleFunc("/signout", signoutGetHandler).Methods("GET")
	r.HandleFunc("/forgot-password", forgotpasswordGetHandler).Methods("GET")
}

// Handlers

func signupGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-signup.html", nil)
}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signin", 302)
}

func signinGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signin.html", nil)
}

func signinPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func signoutGetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}
func forgotpasswordGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-forgotpassword.html", nil)
}
