// cmd/server/handlers.go
// handlers.go contains all request handling logic

package main

import (
	"encoding/json" // For working with JSON data
	"net/http" // for web server functionality
	"time" //For working with dates/times
)

// REQUEST STRUCTURE

// ShippingRequest represents our API response


// TODO MAKE HANDLER ROBUST - KEEP IT SIMPLE
type HomeRequest struct{
	UserName string `json:"username"`
	MessageType string `json:"message_type"`
	Content string `json:"content"`
	Email string `json:"email"`
	RequestID string `json:"requesst_id"`
	Timestamp string `json:"timestamp"`
}



type ShippingRequest struct {
	FromZipCode string  `json:"from_zip"` // where package STARTS
	ToZipCode   string  `json:"to_zip"`   // Where package goes
	Weight      float64 `json:"weight"`   // package weight (can have decimals)
	PackageType string  `json:"package_type"`   //type of package
}

// RESPONSE STRUCTURE

type ShippingResponse struct {
	EstimatedCost float64   `json:"estimated_cost"` // Price (with decimals)
	DeliveryDays  int       `json:"delivery_days"` // Whole number of days
	ServiceType   string    `json:"service_type"` // Type of shipping service
	Timestamp     time.Time `json:"timestamp"`   // When calculation was made
}

// handleHome is our root endpoint handler
// w : where we write our response to the user
// r : contains all information aboute the incoming request
func handleHome(w http.ResponseWriter, r *http.Request) {
	 // Check if request method is GET
	if r.Method != http.MethodGet {  // If it's not GET(like POST/PUT/DELETE), send an error
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return  // Stop processing if wrong method
	}

	 // Tell the browser/client we're sending JSON data
	w.Header().Set("Content-Type","application/json")

	 // Set JSON content type header before writing any response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Thank you, Welcome to Smail🐌📮 - Your friendly optimization journey begins",
		"version": "1.0",        // The following 3 are our key-value pairs that will become JSON
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
	// Try to read the JSON from the request
	var req ShippingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
		// Validate request data
		// Make sure zip codes aren't empty
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
