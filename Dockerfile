# Usar una imagen base de Go para compilar la aplicación
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el contenido del directorio actual al directorio de trabajo en el contenedor
COPY . .

# Compilar el código de Go
RUN go build -o main .

# Exponer el puerto en el que tu aplicación escucha
EXPOSE 8080

# Comando para ejecutar tu aplicación cuando el contenedor se inicie
CMD ["./main"]