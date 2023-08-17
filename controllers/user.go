package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"freehzaix.com/crud_golang/databases"
	"freehzaix.com/crud_golang/models"
)

func AddUser(w http.ResponseWriter, r*http.Request){
	tmpl := template.Must(template.ParseFiles("views/addUser.html"))
	tmpl.Execute(w, nil)
}

func AddUserPost(w http.ResponseWriter, r*http.Request){
	tmpl := template.Must(template.ParseFiles("views/addUserPost.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	users := models.User{
		Nom:   r.FormValue("nom"),
		Prenom: r.FormValue("prenom"),
		Telephone: r.FormValue("telephone"),
		Email: r.FormValue("email"),
	}

	databases.InsertUser(users)

	tmpl.Execute(w, users)
}

func UpdateUser(w http.ResponseWriter, r*http.Request){
	tmpl := template.Must(template.ParseFiles("views/showUser.html"))

	// Récupérer les paramètres de requête de l'URL
	queryParams := r.URL.Query()

	// Récupérer la valeur du paramètre "id"
	id := queryParams.Get("id")

	id_c, _ := strconv.Atoi(id)

	user := databases.ShowUser(id_c)

	tmpl.Execute(w, user)
}

func UpdateUserPost(w http.ResponseWriter, r*http.Request){

	if r.Method != http.MethodPost {
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	users := models.User{
		Id: id,
		Nom:   r.FormValue("nom"),
		Prenom: r.FormValue("prenom"),
		Telephone: r.FormValue("telephone"),
		Email: r.FormValue("email"),
	}

	databases.UpdateUser(users)

	// Effectuer la redirection vers une autre route
	http.Redirect(w, r, "/", http.StatusFound)
	
}

func DeleteUser(w http.ResponseWriter, r*http.Request){
	
	tmpl2 := template.Must(template.ParseFiles("views/ShowDeleteUser.html"))
	fmt.Println(tmpl2)
	fmt.Println(r.URL.Query())
	
	// Récupérer les paramètres de requête de l'URL
	queryParams := r.URL.Query()
	fmt.Println(queryParams)

	// Récupérer la valeur du paramètre "id"
	id := queryParams.Get("id")

	id_d, _ := strconv.Atoi(id)

	user := databases.ShowUserId(id_d)

	tmpl2.Execute(w, user)
	
}

func DeleteUserPost(w http.ResponseWriter, r*http.Request){

	if r.Method != http.MethodPost {
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	users := models.User{
		Id: id,
	}

	databases.DeleteUser(users)

	// Effectuer la redirection vers une autre route
	http.Redirect(w, r, "/", http.StatusFound)

}