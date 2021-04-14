package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product Type (Object Type)
type Product struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}

var product []Product

// Function GetProducts
func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(product)
}

// Main Function
func main() {
	product = append(product, Product{ID: "1", Name: "Product 1", Price: 3.15})

	router := mux.NewRouter()
	router.HandleFunc("/v1/products", GetProducts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
