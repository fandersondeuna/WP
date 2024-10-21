package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // Ensure the request body is closed after reading

	// Log the raw request body
	log.Printf("Received notification: %s", string(body))

	// Response acknowledgment
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[OK]")) // Acknowledgment response
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/notifications", notificationHandler) // Set up the endpoint
	log.Printf("Server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
