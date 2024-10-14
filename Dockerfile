# Usa una imagen de Go como base
FROM golang:1.23.2

# Establece el directorio de trabajo
WORKDIR /app

# Copia el módulo y los archivos go
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente
COPY . .

# Compila la aplicación
RUN go build -o notifications-server

# Expone el puerto que usará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./notifications-server"]
