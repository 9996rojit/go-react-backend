package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var HandleAuth = func(r *mux.Router) {
	// r.HandleFunc("/", controllers.HandleHome).Methods("GET")
	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	r.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
	//		r.HandleFunc("/forget-password", controllers.HandleForgetPassword).Methods("POST")
	//		r.HandleFunc("/change-password", controllers.HandleChangePassword).Methods("POST")
	//	}
}
