# Weather Service

Serviço em Go que retorna a temperatura atual baseada em um CEP brasileiro.

## Requisitos

- Go 1.21 ou superior
- Docker (opcional)
- Conta no WeatherAPI (https://www.weatherapi.com/)

## Configuração

1. Clone o repositório
2. Copie o arquivo `.env.example` para `.env`:
```bash
cp .env.example .env
```
3. Edite o arquivo `.env` e adicione sua chave da API do WeatherAPI:
```
WEATHER_API_KEY=sua_chave_aqui
```

## Executando localmente

```bash
go run cmd/main.go
```

## Executando com Docker

```bash
docker build -t weather-service .
docker run -p 8080:8080 --env-file .env weather-service
```

## Endpoints

### GET /weather/:zipcode

Retorna a temperatura atual para um CEP específico.

#### Parâmetros

- `zipcode`: CEP brasileiro (8 dígitos)

#### Respostas

Sucesso (200):
```json
{
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
}
```

CEP inválido (422):
```json
{
    "error": "invalid zipcode"
}
```

CEP não encontrado (404):
```json
{
    "error": "can not find zipcode"
}
```

## Deploy no Google Cloud Run

1. Configure o Google Cloud SDK
2. Execute os comandos:
```bash
gcloud builds submit --tag gcr.io/seu-projeto/weather-service
gcloud run deploy weather-service --image gcr.io/seu-projeto/weather-service --platform managed
```

## Testes

```bash
go test ./...
``` 