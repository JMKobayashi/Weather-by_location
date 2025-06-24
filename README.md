# Weather Service

Serviço em Go que retorna a temperatura atual baseada em um CEP brasileiro. O sistema recebe um CEP válido de 8 dígitos, identifica a cidade e retorna o clima atual em Celsius, Fahrenheit e Kelvin.

## 🚀 Funcionalidades

- ✅ Validação de CEP (8 dígitos)
- ✅ Consulta de localização via ViaCEP API
- ✅ Consulta de temperatura via WeatherAPI
- ✅ Conversão automática de temperaturas (Celsius, Fahrenheit, Kelvin)
- ✅ Tratamento adequado de erros
- ✅ Deploy automatizado no Google Cloud Run
- ✅ Testes automatizados

## 📋 Requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (para testes locais)
- Conta no Google Cloud Platform
- Conta no WeatherAPI (https://www.weatherapi.com/) - Plano gratuito disponível

## 🛠️ Configuração

1. Clone o repositório
2. Copie o arquivo de exemplo de variáveis de ambiente:
```bash
cp env.example .env
```

3. Edite o arquivo `.env` e adicione sua chave da API do WeatherAPI:
```bash
WEATHER_API_KEY=sua_chave_aqui
PORT=8080
```

## 🏃‍♂️ Executando localmente

### Com Go diretamente:
```bash
go run cmd/main.go
```

### Com Docker:
```bash
docker build -t weather-service .
docker run -p 8080:8080 --env-file .env weather-service
```

### Com Docker Compose:
```bash
docker-compose up --build
```

## 🧪 Testes

### Executar todos os testes:
```bash
go test ./...
```

### Executar testes com cobertura:
```bash
go test -cover ./...
```

### Executar testes de integração:
```bash
go test -v ./api_test.go
```

## 📡 Endpoints

### GET /health
Endpoint de saúde do serviço.

**Resposta (200):**
```json
{
    "status": "ok"
}
```

### GET /weather/:zipcode
Retorna a temperatura atual para um CEP específico.

**Parâmetros:**
- `zipcode`: CEP brasileiro (8 dígitos, com ou sem hífen)

**Exemplos de uso:**
```bash
curl http://localhost:8080/weather/01310900
curl http://localhost:8080/weather/01310-900
```

**Respostas:**

✅ Sucesso (200):
```json
{
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
}
```

❌ CEP inválido (422):
```json
{
    "error": "invalid zipcode"
}
```

❌ CEP não encontrado (404):
```json
{
    "error": "can not find zipcode"
}
```

## 🚀 Deploy no Google Cloud Run

### Pré-requisitos:
1. Google Cloud SDK instalado e configurado
2. Projeto criado no Google Cloud Platform
3. APIs habilitadas (Cloud Build, Cloud Run, Container Registry)

### Deploy automático:
```bash
./deploy.sh <PROJECT_ID> <WEATHER_API_KEY>
```

### Deploy manual:
```bash
# Configurar projeto
gcloud config set project <PROJECT_ID>

# Build e push da imagem
docker build -t gcr.io/<PROJECT_ID>/weather-service .
docker push gcr.io/<PROJECT_ID>/weather-service

# Deploy no Cloud Run
gcloud run deploy weather-service \
    --image gcr.io/<PROJECT_ID>/weather-service \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --port 8080 \
    --memory 512Mi \
    --cpu 1 \
    --max-instances 10 \
    --set-env-vars WEATHER_API_KEY=<WEATHER_API_KEY>
```

## 📊 APIs Utilizadas

### ViaCEP API
- **URL:** https://viacep.com.br/
- **Uso:** Consulta de localização por CEP
- **Limite:** Gratuito, sem autenticação

### WeatherAPI
- **URL:** https://www.weatherapi.com/
- **Uso:** Consulta de temperatura por localização
- **Limite:** 1.000.000 requests/mês no plano gratuito
- **Autenticação:** API Key obrigatória

## 🔧 Fórmulas de Conversão

- **Celsius para Fahrenheit:** F = C × 1.8 + 32
- **Celsius para Kelvin:** K = C + 273

## 📁 Estrutura do Projeto

```
weather-service/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/
│   │   └── weather_handler.go    # Handlers HTTP
│   ├── models/
│   │   └── weather.go            # Estruturas de dados
│   └── services/
│       ├── weather_service.go    # Lógica de negócio
│       └── weather_service_test.go # Testes unitários
├── Dockerfile               # Configuração Docker
├── docker-compose.yml       # Configuração Docker Compose
├── cloudbuild.yaml          # Configuração Cloud Build
├── deploy.sh               # Script de deploy
├── api_test.go             # Testes de integração
├── api.http                # Exemplos de requisições HTTP
├── env.example             # Exemplo de variáveis de ambiente
└── README.md               # Documentação principal
```

## 📚 Documentação Adicional

Para informações mais detalhadas, consulte os seguintes arquivos:

### 📋 [DEPLOY_INSTRUCTIONS.md](./DEPLOY_INSTRUCTIONS.md)
Instruções detalhadas para deploy no Google Cloud Run, incluindo:
- Configuração inicial do Google Cloud SDK
- Habilitar APIs necessárias
- Deploy automático e manual
- Troubleshooting comum
- Monitoramento e logs
- Informações sobre custos e segurança

### 📊 [IMPLEMENTATION_SUMMARY.md](./IMPLEMENTATION_SUMMARY.md)
Resumo técnico da implementação, incluindo:
- Arquitetura do sistema
- Funcionalidades implementadas
- Cobertura de testes
- Configurações de produção
- APIs externas utilizadas
- Requisitos atendidos

### 📦 [ENTREGA.md](./ENTREGA.md)
Checklist completo de entrega, incluindo:
- Status de todos os requisitos
- Arquivos entregues
- Como testar localmente e no Cloud Run
- Próximos passos para deploy
- Comandos úteis para monitoramento

## 🐛 Troubleshooting

### Erro: "WEATHER_API_KEY environment variable is required"
- Verifique se o arquivo `.env` existe e contém a chave da API
- Verifique se a variável está sendo carregada corretamente

### Erro: "can not find zipcode"
- Verifique se o CEP existe no Brasil
- Teste com um CEP conhecido como "01310900" (São Paulo)

### Erro no Cloud Run: "Container failed to start"
- Verifique se a porta 8080 está sendo exposta
- Verifique se a variável WEATHER_API_KEY está configurada

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.
