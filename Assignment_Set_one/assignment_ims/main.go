package main

import (
	"assignment_ims/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/product/add", handlers.AddProductHandler)
	http.HandleFunc("/product/update-stock", handlers.UpdateStockHandler)
	http.HandleFunc("/product/search", handlers.SearchProductHandler)
	http.HandleFunc("/product/display", handlers.DisplayInventoryHandler)
	http.HandleFunc("/product/sort", handlers.SortProductsHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
