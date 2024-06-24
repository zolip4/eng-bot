package main

import (
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		http.Redirect(w, r, "/welcome?name="+name, http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/form.html"))
	tmpl.Execute(w, nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	tmpl := template.Must(template.ParseFiles("web/templates/welcome.html"))
	tmpl.Execute(w, map[string]string{"Name": name})
}
