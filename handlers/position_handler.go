package handlers

import (
    "go-cv-matcher/models"
    "net/http"
    "strings"
    "text/template"
)

var position models.Position

func ShowPositionForm(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/position.html"))
    tmpl.Execute(w, nil)
}

func SavePosition(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        position.Title = r.FormValue("title")
        skills := r.FormValue("skills")
        position.Skills = strings.Split(skills, ",")
        http.Redirect(w, r, "/upload", http.StatusSeeOther)
    }
}
