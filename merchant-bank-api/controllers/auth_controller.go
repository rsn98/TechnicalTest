package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rsn98/merchant-bank-api/repositories"
)

// Struktur untuk menerima data login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Fungsi login pelanggan
func Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	customer, err := repositories.FindCustomerByEmail(request.Email)
	if err != nil || customer.Password != request.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	customer.IsLoggedIn = true

	// Update status pelanggan ke JSON
	customers, _ := repositories.GetCustomers()
	for i, c := range customers {
		if c.Email == customer.Email {
			customers[i] = *customer
			break
		}
	}
	repositories.SaveCustomers(customers)

	// Kirim respons sukses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}
