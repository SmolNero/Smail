// cmd/server/handlers.go
// handlers.go contains all request handling logic

package main

import (
	"encoding/json" // For working with JSON data - converts Go structs to/from
	"net/http" // for web server functionality - powers the web server and HTTP handling
	"time" //For working with dates/times - adds timestamps to responses
	"strings" //Provides string manipulation functions
	"log" // Works with datess and times  - Logs messages for debugging or monitoring
	"github.com/SmolNero/Smail/pkg/models"
)



// CONSTANTS for security
// Prevents excessive data in requests
// Enforces validation rules (e.g., username length, request size)
const  (
	MaxRequestSize = 1024 * 1024 // 1MB max request size - This prevents someone from sending a massive JSON payload that could crash server
	MaxUsernameLen = 50 // Username cant be longer than 50 characters - prevents massive usernames that could mess up UI
	MinUsernameLen = 3 // be at least 3 char - prevents empty or tiny usernames that could be used
	MaxContentLen = 1000 // 1000 chars is roughly 200 words - prevents memory issues from huge messages
	MinRequestIDLen = 8  // IDs must be at least 8 characters - helps maintain uniqueness in tracking requests
 )

// ShippingRequest represents our API response
// Represents incoming data for the "home" endpoint
// Defines the shape of expected JSON input for validation
type HomeRequest struct{
	// Tag tells Go how to convert this field when working with JSON
	UserName string `json:"username"` /// User ident storage- Will be validated against our MinUsernameLen and MaxUsernameLen consts
	MessageType string `json:"message_type"` //Allows to handle diff requests differently (i.e. greet, update, feedback) - this will be validated agains a list of allowed ,essages
	Content string `json:"content"` // Main message content - stores messsage text :: con -> MaxContentLen
	Email string `json:"email"`    // Contact information :: Must contain @ - needs special validation email
	RequestID string `json:"requesst_id"` // FOR REQ TRCKNG - MUST-> MinRequestIDLen req 
	Timestamp string `json:"timestamp"`  // REC REQS - helps track timing and patterns
}


// Represents data for shipping calculations
// Ensures incoming shipping request data is structured correctly
type ShippingRequest struct {
	FromZipCode string  `json:"from_zip"` // where package STARTS
	ToZipCode   string  `json:"to_zip"`   // Where package goes
	Weight      float64 `json:"weight"`   // package weight (can have decimals)
	PackageType string  `json:"package_type"`   //type of package
}

// c STRUCTURE
type ShippingResponse struct {
	EstimatedCost float64   `json:"estimated_cost"` // Price (with decimals)
	DeliveryDays  int       `json:"delivery_days"` // Whole number of days
	ServiceType   string    `json:"service_type"` // Type of shipping service
	Timestamp     time.Time `json:"timestamp"`   // When calculation was made
}

// validateContentType ensures JSON content type - sercurity function that checks IDs
func validateContentType(r *http.Request) bool{ // Declares our functions
	contentType := r.Header.Get("Content-Type") // Gets the content-Types from the request headers
	return strings.Contains(contentType, "application/json") // Checks if we have JSON content
}

// validateHomeRequest validates home endpoint requests
// Consider this function as a form validator
// Func takes a HomeRequest and returns a slice of strings (for error messages)
func validateHomeRequest(req HomeRequest) []string{  
	var errors []string // initializes an empty slice of collect errors

	// username validation
	if req.UserName == "" {		// IF username is empty
		errors = append(errors, "username is required")
	} else if len(req.UserName) < MinUsernameLen {		// ELSE Meets min (3 characters)
		errors = append(errors, "username must be at least 3 characters")
	} else if len(req.UserName) > MaxUsernameLen {		// Does not exceed length (50 characters)
		errors = append(errors, "useername must not exceed 50 characters")
	}

	//MessageType validation
	// Only validates if MessageType is provided
	if req.MessageType != "" {
		validTypes := []string{"greeting", "update", "feedback"}  // Checks against allowed types: "greeting", "update", "Feedback'
		isValid := false
		for _, t := range validTypes {	// Loops through them to check if the provided type matches
			if req.MessageType == t {
				isValid = true
				break
			}
		}
		if !isValid {	// Adds error if the type isn't in our allowed list - isValid = false
			errors = append(errors, "message_type must be one of: greeting, update, feedback")
		}
	}

	// content validation
	// Length check to prevent overly long messages - MaxContentLen as out limit
	if len(req.Content) > MaxContentLen {
		errors = append(errors, "content must not exceed 1000 characters") 
	}

	// Email validation (basic)
	if req.Email != "" {
		if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, "."){ // checks for the pressence pf @ - checks for pressence of '.'
			errors = append(errors, "invalid email format") // adds an error if either is missing

	}
}

	// RequestID Validation
	if req.RequestID == "" {		// Ensures RequestID is provided
		errors = append(errors, "request_id is required")
	} else if len(req.RequestID) < MinRequestIDLen { 	// Checks min length requirements
		errors = append(errors, "request_id must be at least 8 characters")
	}

	return errors	// Returns all collected errors (if any)
}

// handleHome is our root endpoint handler
// w : is like a piece of paper where we'll write our response
// r : contains all the information about the visitor's request
func handleHome(w http.ResponseWriter, r *http.Request) { 	// Consider this func handles visitors to our websites /home 
	w.Header().Set("Content-Type", "application/json")	  // Tells the browser what kind of response we re sending
	// log the request
	log.Printf("Received %s request from %s", r.Method, r.RemoteAddr) // What request we received (GET, POST, etc) :: Who sent it (IP ADDRESS)
	
	switch r.Method {	// Checks the type of request the visitor made (like if viewing or submitting form) - multiway fork
	case http.MethodGet:	// for GET requests (when someone just visits the page)
	 // Set JSON content type header before writing any response
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Thank you, Welcome to Smail🐌📮 - Your friendly optimization journey begins",
			"version": "1.0",        // The following 3 are our key-value pairs that will become JSON
			"status":  "healthy",
	})

	case http.MethodPost, http.MethodPut:	// For POST or PUT requests (when someone submits data)
		// validate content type
		if !validateContentType(r) {	// Checks if we re getting JSON data - if NOT, it will be REJECTED
			http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
			return
		}

		//check request size
		// This makes sure the request isn't too big, like refusing a package that's too heavy.
		if r.ContentLength > MaxRequestSize {
			http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
			return
		}

		var req HomeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Validate request
		errors := validateHomeRequest(req)
		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "errors",				
				"errors": errors,
			})
			return
		}

		//Process valid request
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":"Success",
			"method": r.Method,
			"username": req.UserName,
			"message_type": req.MessageType,
			"email": req.RequestID,
			"received_at": time.Now(),
		})


	default:
		log.Printf("Blocked %s request from %s", r.Method, r.RemoteAddr)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)


	}
}

// func handleShippingCalculate calculates shipping costs
func handleShippingCalculate(w http.ResponseWriter, r *http.Request) {
	 // Set JSON content type
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Received shipping calculation request %s", r.RemoteAddr)

	if r.Method != http.MethodPost {
		log.Printf("Blocked %s request for shipping calculation from %x", r.Method, r.RemoteAddr)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// validate content type
	if !validateContentType(r) {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}

	//Check request size
	if r.ContentLength > MaxRequestSize {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
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

	//Send response
	json.NewEncoder(w).Encode(response)

}

	
	func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
			"time":   time.Now().String(),
	})
}