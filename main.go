package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func main() {
	port := os.Getenv("PORT") // Fly.io establece el puerto en la variable de entorno PORT
	if port == "" {
		port = "8080" // Usar 8080 como puerto por defecto
	}
	http.HandleFunc("/", handler)
	log.Printf("Server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
