package main

import (
	"net/http"
	"freehzaix.com/crud_golang/routes"
)

func main(){

	//Afficher les routes avec la fonction Web()
	routes.Web()

	//Demarrer le serveur sur le port 8080
	http.ListenAndServe(":8080", nil)

}