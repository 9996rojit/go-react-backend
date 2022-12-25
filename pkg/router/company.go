package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var CompanyRoutes = func(r *mux.Router) {
	r.HandleFunc("/companies", controllers.CreateCompany).Methods("POST")
	// r.HandleFunc("/companies/{companyID}", controllers.companyHandler).Methods("GET")
	// r.HandleFunc("/companies/{companyId}", controllers.companyHandler).Methods("PUT")

}
