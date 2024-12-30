package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// ClimateData struct to represent city climate data
type ClimateData struct {
	City         string  `json:"city"`
	AvgTemp      float64 `json:"avgTemp"`
	Rainfall     float64 `json:"rainfall"`
}

// Slice to store climate data
var climateData []ClimateData
var mu sync.Mutex // Mutex for thread-safe access

// AddCityClimateHandler handles adding new city climate data
func AddCityClimateHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newCity ClimateData
	err := json.NewDecoder(r.Body).Decode(&newCity)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	climateData = append(climateData, newCity)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCity)
}

// GetHighestTempCityHandler handles finding the city with the highest temperature
func GetHighestTempCityHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(climateData) == 0 {
		http.Error(w, "No data available", http.StatusNotFound)
		return
	}

	highestCity := climateData[0]
	for _, city := range climateData {
		if city.AvgTemp > highestCity.AvgTemp {
			highestCity = city
		}
	}

	json.NewEncoder(w).Encode(highestCity)
}

// GetLowestTempCityHandler handles finding the city with the lowest temperature
func GetLowestTempCityHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(climateData) == 0 {
		http.Error(w, "No data available", http.StatusNotFound)
		return
	}

	lowestCity := climateData[0]
	for _, city := range climateData {
		if city.AvgTemp < lowestCity.AvgTemp {
			lowestCity = city
		}
	}

	json.NewEncoder(w).Encode(lowestCity)
}

// CalculateAverageRainfallHandler calculates the average rainfall across all cities
func CalculateAverageRainfallHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(climateData) == 0 {
		http.Error(w, "No data available", http.StatusNotFound)
		return
	}

	totalRainfall := 0.0
	for _, city := range climateData {
		totalRainfall += city.Rainfall
	}

	avgRainfall := totalRainfall / float64(len(climateData))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"averageRainfall": avgRainfall})
}

// FilterCitiesByRainfallHandler filters cities by rainfall above a threshold
func FilterCitiesByRainfallHandler(w http.ResponseWriter, r *http.Request) {
	thresholdStr := r.URL.Query().Get("threshold")
	if thresholdStr == "" {
		http.Error(w, "No threshold provided", http.StatusBadRequest)
		return
	}

	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		http.Error(w, "Invalid threshold value", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	var filteredCities []ClimateData
	for _, city := range climateData {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}

	if len(filteredCities) == 0 {
		http.Error(w, "No cities found with rainfall above threshold", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(filteredCities)
}

// SearchCityByNameHandler allows users to search for a city by name
func SearchCityByNameHandler(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		http.Error(w, "No city name provided", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, city := range climateData {
		if strings.EqualFold(city.City, cityName) {
			json.NewEncoder(w).Encode(city)
			return
		}
	}

	http.Error(w, "City not found", http.StatusNotFound)
}
