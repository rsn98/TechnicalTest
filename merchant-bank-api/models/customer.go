package models

type Customer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"is_logged_in"`
}
