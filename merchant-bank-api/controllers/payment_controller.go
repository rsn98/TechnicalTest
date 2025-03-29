package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rsn98/merchant-bank-api/models"
	"github.com/rsn98/merchant-bank-api/repositories"
)

// Struktur untuk menerima request payment
type PaymentRequest struct {
	CustomerID int     `json:"customer_id"`
	MerchantID int     `json:"merchant_id"`
	Amount     float64 `json:"amount"`
}

// Fungsi untuk menangani pembayaran
func Payment(w http.ResponseWriter, r *http.Request) {
	var request PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Cek apakah pelanggan sudah login
	customers, _ := repositories.GetCustomers()
	var customer *models.Customer
	for i, c := range customers {
		if c.ID == request.CustomerID {
			if !c.IsLoggedIn {
				http.Error(w, "Customer is not logged in", http.StatusUnauthorized)
				return
			}
			customer = &customers[i]
			break
		}
	}
	if customer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Cek apakah merchant terdaftar
	merchants, _ := repositories.GetMerchants()
	var merchant *models.Merchant
	for _, m := range merchants {
		if m.ID == request.MerchantID {
			merchant = &m
			break
		}
	}
	if merchant == nil {
		http.Error(w, "Merchant not found", http.StatusNotFound)
		return
	}

	// Simpan transaksi ke history.json
	transaction := models.Transaction{
		CustomerID: request.CustomerID,
		MerchantID: request.MerchantID,
		Amount:     request.Amount,
	}

	err = repositories.SaveTransaction(transaction)
	if err != nil {
		http.Error(w, "Failed to process transaction", http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}
