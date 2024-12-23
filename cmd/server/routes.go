// cmd/server/routes.go

package main

import (
	"net/http"
)

func setupRoutes() {

http.HandleFunc("/", handleHome)
http.HandleFunc("/api/shipping/calculate", handleShippingsCalculate)
http.HandleFunc("/health", handleHealthCheck)	
}