package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TipJarContract struct {
	// Add the fields and methods from your TipJarContract here
}

func main() {
	// Create an instance of your TipJarContract
	tipJar := NewTipJarContract("content_creator_public_key", 30, 0)

	// Define an HTTP route to deposit tips
	http.HandleFunc("/deposit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request body to get user public key and amount
		var depositRequest struct {
			UserPublicKey string `json:"user_public_key"`
			Amount        uint64 `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&depositRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Deposit the tips
		tipJar.Deposit(depositRequest.UserPublicKey, depositRequest.Amount)

		// Return a success response
		w.WriteHeader(http.StatusOK)
	})

	// Define an HTTP route to distribute tips
	http.HandleFunc("/distribute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request body to get the creator's public key
		var distributeRequest struct {
			CreatorPublicKey string `json:"creator_public_key"`
		}

		if err := json.NewDecoder(r.Body).Decode(&distributeRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Distribute the tips (only the creator can do this)
		tipJar.DistributeTips(distributeRequest.CreatorPublicKey)

		// Return a success response
		w.WriteHeader(http.StatusOK)
	})

	// Start the HTTP server
	port := 8080
	fmt.Printf("Listening on :%d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error:", err)
	}
}
