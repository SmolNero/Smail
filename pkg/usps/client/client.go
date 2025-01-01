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
