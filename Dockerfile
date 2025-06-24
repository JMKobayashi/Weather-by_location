# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather-service ./cmd/main.go

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/weather-service .

# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./weather-service"] 