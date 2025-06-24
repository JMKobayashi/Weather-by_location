#!/bin/bash

# Script para deploy no Google Cloud Run
# Uso: ./deploy.sh [PROJECT_ID] [WEATHER_API_KEY]

set -e

# Verificar se as variáveis foram fornecidas
if [ $# -lt 2 ]; then
    echo "Uso: $0 <PROJECT_ID> <WEATHER_API_KEY>"
    echo "Exemplo: $0 meu-projeto minha-chave-api"
    exit 1
fi

PROJECT_ID=$1
WEATHER_API_KEY=$2
SERVICE_NAME="weather-service"
REGION="us-central1"

echo "🚀 Iniciando deploy do Weather Service..."
echo "Projeto: $PROJECT_ID"
echo "Região: $REGION"
echo "Serviço: $SERVICE_NAME"

# Configurar o projeto
echo "📋 Configurando projeto..."
gcloud config set project $PROJECT_ID

# Habilitar APIs necessárias
echo "🔧 Habilitando APIs..."
gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com

# Build e push da imagem
echo "🏗️  Fazendo build da imagem..."
docker build -t gcr.io/$PROJECT_ID/$SERVICE_NAME .

echo "📤 Fazendo push da imagem..."
docker push gcr.io/$PROJECT_ID/$SERVICE_NAME

# Deploy no Cloud Run
echo "🚀 Fazendo deploy no Cloud Run..."
gcloud run deploy $SERVICE_NAME \
    --image gcr.io/$PROJECT_ID/$SERVICE_NAME \
    --platform managed \
    --region $REGION \
    --allow-unauthenticated \
    --port 8080 \
    --memory 512Mi \
    --cpu 1 \
    --max-instances 10 \
    --set-env-vars WEATHER_API_KEY=$WEATHER_API_KEY

# Obter a URL do serviço
SERVICE_URL=$(gcloud run services describe $SERVICE_NAME --platform managed --region $REGION --format 'value(status.url)')

echo "✅ Deploy concluído com sucesso!"
echo "🌐 URL do serviço: $SERVICE_URL"
echo ""
echo "📝 Exemplos de uso:"
echo "curl $SERVICE_URL/health"
echo "curl $SERVICE_URL/weather/01310900"
echo ""
echo "🔧 Para ver os logs:"
echo "gcloud logs tail --service=$SERVICE_NAME --region=$REGION" 