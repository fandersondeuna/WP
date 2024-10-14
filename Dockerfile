# Usa la imagen oficial de Go
FROM golang:1.20

# Establece el directorio de trabajo
WORKDIR /app

# Copia el código Go al contenedor
COPY . .

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que tu aplicación se ejecutará
EXPOSE 8080

# Comando para iniciar la aplicación
CMD ["./main"]
