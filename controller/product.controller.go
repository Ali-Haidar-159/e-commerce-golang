package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"ali.com/market/utils"
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

	utils.CorsHandler(w)
	utils.PreflightHandler(w, r)

	if r.Method != http.MethodGet {
		return
	}

	products, err := os.ReadFile("./data/products.json")
	if err != nil {
		fmt.Println("Error reading the products data:", err)
		http.Error(w, "Could not read products file", http.StatusInternalServerError)
		return
	}

	errJson := json.Unmarshal(products, &list)
	if errJson != nil {
		fmt.Println("Error converting JSON:", err)
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(w, list.Products, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	utils.CorsHandler(w)
	utils.PreflightHandler(w, r)

	if r.Method != http.MethodPost {
		return
	}

	var product1 Product
	var allProducts ProductList

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := os.ReadFile("./data/products.json")
	err1 := json.Unmarshal(data, &allProducts)
	if err1 != nil {
		http.Error(w, "Find error to unmarshal data", http.StatusInternalServerError)
		return
	}

	list := allProducts.Products
	list = append(list, product1)

	utils.SendJsonResponse(w, list, http.StatusCreated)
}
