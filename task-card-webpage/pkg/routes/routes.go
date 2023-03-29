package routes

import (
	"net/http"

	"github.com/luuisavelino/github.com/luuisavelino/task-card-webpage/pkg/controllers"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	// Users
	http.HandleFunc(apiVersion+"/users/", controllers.Index)
	http.HandleFunc(apiVersion+"/users/new", controllers.New)
	http.HandleFunc(apiVersion+"/users/insert", controllers.Insert)
	http.HandleFunc(apiVersion+"/users/delete", controllers.Delete)
	http.HandleFunc(apiVersion+"/users/edit", controllers.Edit)
	http.HandleFunc(apiVersion+"/users/update", controllers.Update)

	// Cards
	http.HandleFunc(apiVersion+"/cards/", controllers.Index)
	http.HandleFunc(apiVersion+"/cards/new", controllers.New)
	http.HandleFunc(apiVersion+"/cards/insert", controllers.Insert)
	http.HandleFunc(apiVersion+"/cards/delete", controllers.Delete)
	http.HandleFunc(apiVersion+"/cards/edit", controllers.Edit)
	http.HandleFunc(apiVersion+"/cards/update", controllers.Update)
	http.HandleFunc(apiVersion+"/cards/move", nil)

	http.ListenAndServe(":8080", nil)
}
