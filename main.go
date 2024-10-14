package main

import (
	"fmt"
	"net/http"
)

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	// Procesar la notificación
	fmt.Fprintf(w, "Notificación recibida")
}

func main() {
	http.HandleFunc("/notifications", notificationHandler)
	fmt.Println("Server is listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
