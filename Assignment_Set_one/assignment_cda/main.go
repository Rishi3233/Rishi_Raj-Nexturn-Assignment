package main

import (
	"assignment_cda/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/climate/add", handlers.AddCityClimateHandler)
	http.HandleFunc("/climate/highest-temp", handlers.GetHighestTempCityHandler)
	http.HandleFunc("/climate/lowest-temp", handlers.GetLowestTempCityHandler)
	http.HandleFunc("/climate/average-rainfall", handlers.CalculateAverageRainfallHandler)
	http.HandleFunc("/climate/filter-rainfall", handlers.FilterCitiesByRainfallHandler)
	http.HandleFunc("/climate/search", handlers.SearchCityByNameHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
