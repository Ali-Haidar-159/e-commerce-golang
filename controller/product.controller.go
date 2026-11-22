package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

var list ProductList

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", http.StatusOK)
		return
	}

	if r.Method == http.MethodOptions {
		fmt.Println("OPTIONS method is triggered")
	}

	products, err := os.ReadFile("./data/products.json")

	if err != nil {
		fmt.Println("Error reading the products data:", err)
		http.Error(w, "Could not read products file", http.StatusInternalServerError)
		return
	} else {

		errJson := json.Unmarshal(products, &list)
		if errJson != nil {
			fmt.Println("Error converting JSON:", err)
			http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Author", "Ali")
	json.NewEncoder(w).Encode(list.Products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Println("Please send POST request")
		http.Error(w, "Please send POST request", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var product1 Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var allProducts ProductList
	data, _ := os.ReadFile("./data/products.json")
	err1 := json.Unmarshal(data, &allProducts)

	if err1 != nil {
		http.Error(w, "Find error to unmarshal data", http.StatusInternalServerError)
		return
	}

	list := allProducts.Products
	list = append(list, product1)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}
