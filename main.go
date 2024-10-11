package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// Define la estructura para procesar el XML
type Notification struct {
    XMLName xml.Name `xml:"Notification"`
    Status  string   `xml:"Status"`
    // Agrega otros campos según la estructura del XML de Worldpay
}

func main() {
    http.HandleFunc("/notificaciones", handleNotifications)
    log.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleNotifications(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Leer el cuerpo de la solicitud
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error al leer el cuerpo", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    // Procesar el XML
    var notification Notification
    err = xml.Unmarshal(body, &notification)
    if err != nil {
        http.Error(w, "Error al procesar el XML", http.StatusBadRequest)
        return
    }

    // Aquí puedes manejar la notificación (ej. guardarla en la base de datos, etc.)
    fmt.Printf("Notificación recibida: %v\n", notification)

    // Responder a Worldpay
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Notificación recibida")
}
