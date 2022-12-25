package router

import (
	"github.com/9996rojit/backend_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var HandelPermissionRoute = func(router *mux.Router) {
	router.HandleFunc("/permission", controllers.CreatePermission).Methods("POST")
	router.HandleFunc("/getPermissionById/{permissionId}", controllers.GetPermissionById).Methods("GET")
	router.HandleFunc("/getPermission", controllers.GetAllPermission).Methods("GET")
	router.HandleFunc("/deletePermission/{permissionId}", controllers.DeletePermissionById).Methods("DELETE")
}
