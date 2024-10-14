# Usa la imagen de Go
FROM golang:1.20

# Establece el directorio de trabajo
WORKDIR /app

# Copia el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto del código
COPY . .

# Compila el código
RUN go build -o notifications-server

# Expone el puerto que usará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./notifications-server"]
