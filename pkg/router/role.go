package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RoleRouter = func(r *mux.Router) {
	r.HandleFunc("/getRoles", controllers.GetRoles).Methods("GET")
	r.HandleFunc("/createRole", controllers.CreateRole).Methods("POST")
	r.HandleFunc("/deleteRole/{roleId}", controllers.DeleteRoleById).Methods("DELETE")
	// r.HandleFunc("/updateRole", controller.UpdateRole).Methods("PUT")
}
