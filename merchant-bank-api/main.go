package main

import (
	"fmt"
	"net/http"

	"github.com/rsn98/merchant-bank-api/routes"
)

func main() {
	r := routes.RegisterRoutes()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
