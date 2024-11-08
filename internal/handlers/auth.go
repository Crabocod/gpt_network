package handlers

import (
	"html/template"
	"net/http"

	"web.app/internal/models"
	"web.app/internal/sessions"
	"web.app/internal/utils"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "register.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordHash := utils.HashPassword(password)

	err := models.RegisterUser(username, passwordHash)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.GetSession(r)
	userID, ok := session.Values["user_id"]

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "home.html", user)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.GetSession(r)
	delete(session.Values, "user_id")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordHash := utils.HashPassword(password)

	user, err := models.GetUserByUsernameAndPassword(username, passwordHash)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	session, _ := sessions.GetSession(r)
	session.Values["user_id"] = user.ID
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
