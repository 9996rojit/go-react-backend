package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var CategoryRoutes = func(r *mux.Router) {
	r.HandleFunc("/create-category", controllers.CreateCategory).Methods("POST")
	// r.HandleFunc("/get-category", controllers.GetCategory).Methods("GET")

}
