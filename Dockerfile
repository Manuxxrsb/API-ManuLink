# Usar una imagen base de Golang
FROM golang:latest

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente de la aplicación
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Establecer las variables de entorno
ENV PGHost=aws-0-us-west-1.pooler.supabase.com
ENV PGPort=5432
ENV PGUser=postgres.jfqnlhzssvesahswlcla
ENV PGPassword=manulink2024
ENV PGName=postgres

# Exponer el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
