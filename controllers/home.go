package controllers

import (
	"html/template"
	"net/http"

	"freehzaix.com/crud_golang/databases"
)

func Home(w http.ResponseWriter, r*http.Request){
	tmpl := template.Must(template.ParseFiles("views/home.html"))
	var users = databases.ReadUser()
	tmpl.Execute(w, users)
}
