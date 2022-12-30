package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
	"github.com/gorilla/mux"
)

var PermissionModel models.Permission

func CreatePermission(w http.ResponseWriter, r *http.Request) {
	permissionBody := &models.Permission{}
	json.NewDecoder(r.Body).Decode(permissionBody)
	fmt.Println("This is request body:%w", permissionBody)
	per := permissionBody.CreatePermisssion()
	res, err := json.Marshal(per)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Write([]byte("Error While parsing json"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetPermissionById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["permissionId"]
	permissionData, _ := models.GetPermissionById(id)
	res, err := json.Marshal(permissionData)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error While getting data from database "))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetAllPermission(w http.ResponseWriter, r *http.Request) {
	allPermission := models.GetPermissions()
	res, err := json.Marshal(allPermission)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Write([]byte("Error While formatting data to json format"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func DeletePermissionById(w http.ResponseWriter, r *http.Request) {
	var permission *models.Permission
	params := mux.Vars(r)
	id := params["permissionId"]
	permission = models.DeletePermissionById(id)
	res, err := json.Marshal(permission)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error While formatting data to json format"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
