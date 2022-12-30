package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var ProductRoutes = func(router *mux.Router) {
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
}
