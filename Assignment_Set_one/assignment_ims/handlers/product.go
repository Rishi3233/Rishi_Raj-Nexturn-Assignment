package handlers

import (
	"encoding/json"

	"net/http"
	"sort"
	"strconv"
	"sync"
)

// Product struct to represent a product
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Slice to store products
var products []Product
var mu sync.Mutex // Mutex for thread-safe access

// AddProductHandler handles adding new products
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if ID is unique
	for _, prod := range products {
		if prod.ID == newProduct.ID {
			http.Error(w, "Product ID must be unique", http.StatusBadRequest)
			return
		}
	}

	// Validate stock and price
	if newProduct.Stock < 0 {
		http.Error(w, "Stock cannot be negative", http.StatusBadRequest)
		return
	}
	if newProduct.Price < 0 {
		http.Error(w, "Price cannot be negative", http.StatusBadRequest)
		return
	}

	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

// UpdateStockHandler handles stock updates
func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	type StockUpdateRequest struct {
		ID    int `json:"id"`
		Stock int `json:"stock"`
	}

	var req StockUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find the product and update the stock
	for i, prod := range products {
		if prod.ID == req.ID {
			if req.Stock < 0 {
				http.Error(w, "Stock cannot be negative", http.StatusBadRequest)
				return
			}
			products[i].Stock = req.Stock
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// SearchProductHandler handles product search by ID or name
func SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	mu.Lock()
	defer mu.Unlock()

	for _, prod := range products {
		if strconv.Itoa(prod.ID) == query || prod.Name == query {
			json.NewEncoder(w).Encode(prod)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// DisplayInventoryHandler displays all products
func DisplayInventoryHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(products) == 0 {
		http.Error(w, "No products in inventory", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// SortProductsHandler sorts products by price or stock
func SortProductsHandler(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sortBy")

	mu.Lock()
	defer mu.Unlock()

	switch sortBy {
	case "price":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	case "stock":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Stock < products[j].Stock
		})
	default:
		http.Error(w, "Invalid sort parameter", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(products)
}
