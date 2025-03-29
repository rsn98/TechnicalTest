package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/rsn98/merchant-bank-api/models"
)

// Fungsi untuk membaca semua pelanggan dari JSON
func GetCustomers() ([]models.Customer, error) {
	file, err := os.ReadFile("data/customers.json")
	if err != nil {
		return nil, err
	}

	var customers []models.Customer
	err = json.Unmarshal(file, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// Fungsi untuk menyimpan pelanggan ke JSON
func SaveCustomers(customers []models.Customer) error {
	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data/customers.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk mencari pelanggan berdasarkan email
func FindCustomerByEmail(email string) (*models.Customer, error) {
	customers, err := GetCustomers()
	if err != nil {
		return nil, err
	}

	for _, customer := range customers {
		if customer.Email == email {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}

func LogoutCustomer(customerID int) error {
	customers, err := GetCustomers()
	if err != nil {
		return err
	}

	var found bool
	for i, c := range customers {
		if c.ID == customerID {
			if !c.IsLoggedIn {
				return errors.New("customer is already logged out")
			}
			customers[i].IsLoggedIn = false
			found = true
			break
		}
	}

	if !found {
		return errors.New("customer not found")
	}

	// Simpan perubahan ke JSON
	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/customers.json", data, 0644)
}