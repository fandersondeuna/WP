package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type PaymentNotification struct {
	XMLName      xml.Name `xml:"paymentService"`
	MerchantCode string   `xml:"merchantCode,attr"`
	Notify       Notify   `xml:"notify"`
}

type Notify struct {
	OrderStatusEvent OrderStatusEvent `xml:"orderStatusEvent"`
}

type OrderStatusEvent struct {
	OrderCode string  `xml:"orderCode,attr"`
	Payment   Payment `xml:"payment"`
}

type Payment struct {
	PaymentMethod string `xml:"paymentMethod"`
	Amount        Amount `xml:"amount"`
}

type Amount struct {
	Value        int    `xml:"value,attr"`
	CurrencyCode string `xml:"currencyCode,attr"`
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification PaymentNotification
	if err := xml.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received notification for order: %s, amount: %d %s",
		notification.Notify.OrderStatusEvent.OrderCode,
		notification.Notify.OrderStatusEvent.Payment.Amount.Value,
		notification.Notify.OrderStatusEvent.Payment.Amount.CurrencyCode)
}

func main() {
	http.HandleFunc("/notifications", notificationHandler)
	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
