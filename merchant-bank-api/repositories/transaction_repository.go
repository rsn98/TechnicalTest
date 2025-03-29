package repositories

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rsn98/merchant-bank-api/models"
)

// Fungsi untuk membaca transaksi dari history.json
func GetTransactions() ([]models.Transaction, error) {
	file, err := os.ReadFile("data/history.json")
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	err = json.Unmarshal(file, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// Fungsi untuk menyimpan transaksi ke history.json
func SaveTransaction(transaction models.Transaction) error {
	transactions, _ := GetTransactions()

	// Set ID otomatis
	transaction.ID = len(transactions) + 1
	transaction.Timestamp = time.Now().Format(time.RFC3339)

	transactions = append(transactions, transaction)

	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/history.json", data, 0644)
}
