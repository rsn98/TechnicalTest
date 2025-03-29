package repositories

import (
	"encoding/json"
	"os"

	"github.com/rsn98/merchant-bank-api/models"
)

// Fungsi untuk membaca daftar merchant dari JSON
func GetMerchants() ([]models.Merchant, error) {
	file, err := os.ReadFile("data/merchants.json")
	if err != nil {
		return nil, err
	}

	var merchants []models.Merchant
	err = json.Unmarshal(file, &merchants)
	if err != nil {
		return nil, err
	}

	return merchants, nil
}
