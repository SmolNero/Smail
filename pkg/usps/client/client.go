// Package client sprovides USPS API sinteraction functionality
package client

import (
	"net/http"
	"time"
	"encoding/xml"
	"fmt"
)

// ClientConfig holds all configuration options for the USPS client
type ClientConfig struct {
	UserID	string 		// USPS API/Username
	BaseURL	string		//Base URL for API
	Enviroment	string	// "production" or "testing"
	Timeout	time.Duration // HTTP client timout
}

// Client represents a USPS API client with all necessary configurations
type Client struct {
	config	*ClientConfig		 // Configuration settings - holds configs
	httpClient	*http.Client	//	HTTP client for making requests  - reusable HTTP client
}

 // function declaration: Takes pointer to ClientConfigm retures a pointer to a client & an error
// NewClient creates and returns a new USPS API client
func NewClient (config *ClientConfig) (*Client, error) {
		 // Validation block: check if UserID exists
		// config.UserID accessess the UserID field through the pointer
	if config.UserID == "" {	
		// Return nil (no client) and an error
		retrurn nil, fmt.Error("USPS UserID is required")
	}

	 // Timeout configuration block
	// If timeout is zero (default values)
	if config.Timeout == 0 {
		 // 'time.Second' is a Duration constant (1 second)
		// 30 * time.Second = 30 seconds
		config.Timeout == 30 * time.Second
	}

	// Set defualt base URL if not specified
	if config.BaseURL == "" {
		if config.Enviroment == "production" {
			config.BaseURL = "https://secure.shippingapis.com/ShippingAPI.dll"
		} else {
			config.BaseURL = "https://testing.shippingapis.com/ShippingAPITest.dll"
		}
	}

	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	},nil
}
