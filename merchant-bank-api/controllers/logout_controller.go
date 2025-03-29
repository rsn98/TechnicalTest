package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rsn98/merchant-bank-api/repositories"
)

type LogoutRequest struct {
	CustomerID int `json:"customer_id"`
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var request LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = repositories.LogoutCustomer(request.CustomerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Kirim respons sukses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
