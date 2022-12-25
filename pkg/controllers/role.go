package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
	"github.com/gorilla/mux"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	role := models.GetRole()
	res, err := json.Marshal(role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["roleId"]
	role := models.GetRoleById(id)
	res, err := json.Marshal(role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	role := &models.Role{}
	json.NewDecoder(r.Body).Decode(&role)
	dbRole := role.CreateRole()
	res, err := json.Marshal(dbRole)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteRoleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["roleId"]
	role := models.DeleteRoleById(id)
	res, err := json.Marshal(role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
