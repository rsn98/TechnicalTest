package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rsn98/merchant-bank-api/routes"
)

func setupLogging() {
	file, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Logging initialized")
}

func init() {
	setupLogging()
}


func main() {
	r := routes.RegisterRoutes()
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
