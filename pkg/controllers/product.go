package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	Product := &models.Product{}
	json.NewDecoder(r.Body).Decode(Product)
	dbProduct := Product.CreateProduct()

	res, err := json.Marshal(dbProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}
