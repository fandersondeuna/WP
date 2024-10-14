package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Estructuras para deserializar la notificación XML
type PaymentService struct {
	XMLName      xml.Name `xml:"paymentService"`
	Version      string   `xml:"version,attr"`
	MerchantCode string   `xml:"merchantCode,attr"`
	Notify       Notify   `xml:"notify"`
}

type Notify struct {
	OrderStatusEvent OrderStatusEvent `xml:"orderStatusEvent"`
}

type OrderStatusEvent struct {
	OrderCode string  `xml:"orderCode,attr"`
	Payment   Payment `xml:"payment"`
	Journal   Journal `xml:"journal"`
}

type Payment struct {
	PaymentMethod string `xml:"paymentMethod"`
	Amount        Amount `xml:"amount"`
	LastEvent     string `xml:"lastEvent"`
}

type Amount struct {
	Value                string `xml:"value,attr"`
	CurrencyCode         string `xml:"currencyCode,attr"`
	Exponent             string `xml:"exponent,attr"`
	DebitCreditIndicator string `xml:"debitCreditIndicator,attr"`
}

type Journal struct {
	JournalType string `xml:"journalType,attr"`
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	// Lee el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Decodifica el XML de la notificación
	var paymentService PaymentService
	err = xml.Unmarshal(body, &paymentService)
	if err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	// Manejo de la notificación
	log.Printf("[%s] Received notification: XML=%s, order: %s, status: %s, payment method: %s, amount: %s %s, last event: %s, journal type: %s",
		time.Now().Format(time.RFC3339),
		string(body), // Incluye el XML completo
		paymentService.Notify.OrderStatusEvent.OrderCode,
		paymentService.Notify.OrderStatusEvent.Payment.LastEvent,
		paymentService.Notify.OrderStatusEvent.Payment.PaymentMethod,
		paymentService.Notify.OrderStatusEvent.Payment.Amount.Value,
		paymentService.Notify.OrderStatusEvent.Payment.Amount.CurrencyCode,
		paymentService.Notify.OrderStatusEvent.Payment.LastEvent,
		paymentService.Notify.OrderStatusEvent.Journal.JournalType)

	// Respuesta de reconocimiento
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[OK]")) // Acknowledgment response
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/notifications", notificationHandler) // Configura el endpoint
	log.Printf("Server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
