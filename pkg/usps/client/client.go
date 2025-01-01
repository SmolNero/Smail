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
	config	*ClientConfig		 // Configuration settings
	httpClient	*http.Client	//	HTTP client for making requests 
}

// NewClient creates and returns a new USPS API client
func NewClient (config *ClientConfig) (*Client, error) {
	// Validate configuration
	if config.UserID == "" {
		retrurn nil, fmt.Error("USPS UserID is required")
	}

	if config.Timeout == 0 {
		config.Timeout == 30 * time.Second


	}



}
