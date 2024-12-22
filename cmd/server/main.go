// cmd/server/main.go

package main

// This line says "this file is part of our main program"
import (
	"fmt" // printing is for printing things
	"log" // log is for printing important messages
	"net/http" // lets create web servers
)

// This function runs when someone visits our website
// w http.ResponseWriter - This how we send stuff back to the browswer
// r *http.Request - This contains all the info about whos visiting
// w and r are conventional names in Go for response writer and request

func handleHome(w http.ResponseWriter, r *http.Request) {

	// when someone visits, send them a message
	// Fprintf is a formatted print function that writes to a writer interface

	fmt.Fprintf(w, "Thank you, Welcome to Smail🐌📮")
	fmt.Fprintf(w, "Your friendly optimization journey begins")
}

// This is where our program starts

func main() {
	// print message when we start

	fmt.Println("Smail🐌📮 - a friendly service" )

	// When someone visits our website (/) show them the home page

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/shipping/calculate", handleShippingCalculate)
	http.HandleFunc("/health", handleHealthCheck)	

	// Start our web server on port 8080

	port := ":8080"
	fmt.Println("sever is ready! Visit http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


