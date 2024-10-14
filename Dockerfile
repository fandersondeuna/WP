# Usa la imagen de Go oficial
FROM golang:1.20

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto de los archivos de tu aplicación
COPY . .

# Compila el proyecto
RUN go build -o notifications-server .

# Expone el puerto
EXPOSE 8080

# Comando para ejecutar tu aplicación
CMD ["./notifications-server"]
