package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

type Notification struct {
    XMLName        xml.Name `xml:"paymentService"`
    MerchantCode   string   `xml:"merchantCode,attr"`
    Notify         Notify   `xml:"notify"`
}

type Notify struct {
    OrderStatusEvent OrderStatusEvent `xml:"orderStatusEvent"`
}

type OrderStatusEvent struct {
    OrderCode string `xml:"orderCode"`
    Payment   Payment `xml:"payment"`
}

type Payment struct {
    PaymentMethod string  `xml:"paymentMethod"`
    Amount        Amount  `xml:"amount"`
    LastEvent     string  `xml:"lastEvent"`
}

type Amount struct {
    Value         int    `xml:"value,attr"`
    CurrencyCode  string `xml:"currencyCode,attr"`
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // fallback
    }

    http.HandleFunc("/notifications", handleNotifications)

    log.Printf("Server is running on port %s", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}

func handleNotifications(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Read and parse the XML body
    body, err := ioutil.ReadAll(r.Body)
    if
