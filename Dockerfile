FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weather-service ./cmd/main.go

FROM scratch
COPY --from=builder /app/weather-service .
CMD ["./weather-service"] 