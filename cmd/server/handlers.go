package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// ShippingRequest represents our API response
type ShippingRequest struct {
	FromZipCode string  `json:"from_zip"`
	ToZipCode   string  `json:"to_zip"`
	Weight      float64 `json:"weight"`
	PackageType string  `json:"package_type"`
}

type ShippingResponse struct {
	EstimatedCost float64   `json:"estimated_cost"`
	DeliveryDays  int       `json:"delivery_days"`
	ServiceType   string    `json:"service_type"`
	Timestamp     time.Time `json:"timestamp"`
}

// handleHome is our root endpoint handler
func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to Smail API",
		"version": "1.0",
		"status":  "healthy",
	})
}

// func handleShippingCalculate calculates shipping costs
func handleShippingCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Parsings incoming JSON request
	var req ShippingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// Validate request data
	if req.FromZipCode == "" || req.ToZipCode == "" {
		http.Error(w, "Missing zip code", http.StatusBadRequest)
		return
	}

	if req.Weight <= 0 {
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}

	// Create example response (We'll integrate with USPS API later)
	response := ShippingResponse{
		EstimatedCost: 15.99, // Placeholder
		DeliveryDays:  3,     // Placeholder
		ServiceType:   "Priority Mail",
		Timestamp:     time.Now(),
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	//Send response
	json.NewEncoder(w).Encode(response)

}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().String(),
	})
}
