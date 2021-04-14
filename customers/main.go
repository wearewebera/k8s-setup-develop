package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var customer []Customer

// Function GetCustomers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customer)
}

// Main Function
func main() {
	customer = append(customer, Customer{ID: "1", Firstname: "John", Lastname: "Doe"})

	router := mux.NewRouter()
	router.HandleFunc("/v1/customers", GetCustomers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
