package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
	"github.com/gorilla/mux"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	company := &models.Company{}
	CompanyRes := company.CreateCompany()

	res, err := json.Marshal(CompanyRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error While Parsing json"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	company := models.GetCompanyByID(id)
	res, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error While Parsing json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)

}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	company := models.GetCompanyByID(id)
	res, err := json.Marshal(company)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error While Parsing json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
