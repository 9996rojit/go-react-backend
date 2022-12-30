package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := &models.Category{}

	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decode := json.NewDecoder(r.Body).Decode(category)
	fmt.Println("This is body of category: %w", decode)
}
