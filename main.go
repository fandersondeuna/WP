package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Received notification: %s", body)

	// Acknowledge receipt
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[OK]"))
}

func main() {
	http.HandleFunc("/notifications", notificationHandler)
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
