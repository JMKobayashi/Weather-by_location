# 🚀 Instruções de Deploy - Weather Service

## 📋 Pré-requisitos

1. **Google Cloud SDK** instalado e configurado
2. **Docker** instalado e funcionando
3. **Conta no Google Cloud Platform** com projeto criado
4. **Chave da API do WeatherAPI** (gratuita em https://www.weatherapi.com/)

## 🔧 Configuração Inicial

### 1. Configurar o Google Cloud SDK
```bash
# Instalar o Google Cloud SDK (se ainda não tiver)
# https://cloud.google.com/sdk/docs/install

# Fazer login
gcloud auth login

# Configurar o projeto
gcloud config set project SEU_PROJECT_ID
```

### 2. Habilitar APIs necessárias
```bash
gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com
```

### 3. Configurar variáveis de ambiente
```bash
# Copiar o arquivo de exemplo
cp env.example .env

# Editar o arquivo .env e adicionar sua chave da API
echo "WEATHER_API_KEY=sua_chave_aqui" > .env
```

## 🚀 Deploy

### Deploy manual
```bash
# 1. Build da imagem Docker
docker build -t gcr.io/SEU_PROJECT_ID/weather-service .

# 2. Push para o Container Registry
docker push gcr.io/SEU_PROJECT_ID/weather-service

# 3. Deploy no Cloud Run
gcloud run deploy weather-service \
    --image gcr.io/SEU_PROJECT_ID/weather-service \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --port 8080 \
    --memory 512Mi \
    --cpu 1 \
    --max-instances 10 \
    --set-env-vars WEATHER_API_KEY=SUA_WEATHER_API_KEY
```

## 🧪 Testando o Deploy

### 1. Obter a URL do serviço
```bash
gcloud run services describe weather-service \
    --platform managed \
    --region us-central1 \
    --format 'value(status.url)'
```

### 2. Testar os endpoints
```bash
# Teste de saúde
curl https://SEU_SERVICE_URL/health

# Teste com CEP válido
curl https://SEU_SERVICE_URL/weather/01310900

# Teste com CEP inválido
curl https://SEU_SERVICE_URL/weather/1234567

# Teste com CEP inexistente
curl https://SEU_SERVICE_URL/weather/99999999
```

## 📊 Monitoramento

### Ver logs em tempo real
```bash
gcloud logs tail --service=weather-service --region=us-central1
```

### Ver métricas do serviço
```bash
gcloud run services describe weather-service \
    --platform managed \
    --region us-central1
```

## 🔧 Troubleshooting

### Erro: "Container failed to start"
- Verifique se a variável `WEATHER_API_KEY` está configurada
- Verifique se a porta 8080 está sendo exposta no Dockerfile

### Erro: "Permission denied"
- Execute `gcloud auth login` novamente
- Verifique se tem permissões no projeto

### Erro: "API not enabled"
- Execute os comandos para habilitar as APIs necessárias

### Erro: "Image not found"
- Verifique se o build e push da imagem foram bem-sucedidos
- Verifique se o nome da imagem está correto

## 💰 Custos

O Google Cloud Run tem um **tier gratuito generoso**:
- **2 milhões de requests/mês** gratuitos
- **360.000 vCPU-segundos/mês** gratuitos
- **180.000 GiB-segundos/mês** gratuitos

Para este projeto, você provavelmente não pagará nada, pois:
- O serviço é stateless e escala para zero
- As requisições são leves
- O uso típico é baixo

## 🔒 Segurança

### Variáveis de ambiente
- A `WEATHER_API_KEY` é armazenada como variável de ambiente no Cloud Run
- Não é exposta no código ou logs

### Acesso público
- O serviço está configurado como `--allow-unauthenticated`
- Isso permite acesso público sem autenticação
- Para restringir acesso, remova essa flag

## 📈 Escalabilidade

O Cloud Run automaticamente:
- Escala para zero quando não há tráfego
- Escala horizontalmente conforme a demanda
- Gerencia a infraestrutura automaticamente

## 🔄 Atualizações

Para atualizar o serviço:
```bash
# Rebuild e redeploy
docker build -t gcr.io/SEU_PROJECT_ID/weather-service .
docker push gcr.io/SEU_PROJECT_ID/weather-service
gcloud run deploy weather-service --image gcr.io/SEU_PROJECT_ID/weather-service
```

## 📞 Suporte

Se encontrar problemas:
1. Verifique os logs: `gcloud logs tail --service=weather-service`
2. Teste localmente primeiro
3. Verifique se todas as APIs estão habilitadas
4. Consulte a documentação do Google Cloud Run 