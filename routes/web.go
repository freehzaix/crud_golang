package routes

import (
	"net/http"

	"freehzaix.com/crud_golang/controllers"
)

func Web(){
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/AddUser", controllers.AddUser)
	http.HandleFunc("/AddUserPost", controllers.AddUserPost)
	http.HandleFunc("/ShowUser/", controllers.UpdateUser)
	http.HandleFunc("/UpdateUserPost", controllers.UpdateUserPost)
	http.HandleFunc("/ShowDeleteUser/", controllers.DeleteUser)
	http.HandleFunc("/DeleteUserPost", controllers.DeleteUserPost)
}