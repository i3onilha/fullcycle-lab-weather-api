# FullCycle Lab Weather API

Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin).

## 📋 Requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (para testes)
- Conta na [WeatherAPI](https://www.weatherapi.com/) (gratuita)
- Conta no Google Cloud Platform (para deploy no Cloud Run)

## 🚀 Como executar localmente

### Pré-requisitos

1. Obtenha uma API key gratuita em [WeatherAPI](https://www.weatherapi.com/)
2. Configure a variável de ambiente:

```bash
export WEATHER_API_KEY=sua-api-key-aqui
```

### Executando com Go

```bash
# Instalar dependências
go mod download

# Executar
go run main.go
```

O servidor estará disponível em `http://localhost:8080`

### Executando com Docker Compose

```bash
# Configurar a API key
export WEATHER_API_KEY=sua-api-key-aqui

# Executar com docker-compose
docker-compose up --build
```

## 📡 Endpoints

### GET /weather/{cep}

Retorna as temperaturas em Celsius, Fahrenheit e Kelvin para o CEP informado.

**Parâmetros:**
- `cep`: CEP brasileiro com 8 dígitos (pode incluir hífen ou não)

**Exemplo de requisição:**
```bash
curl http://localhost:8080/weather/01310100
```

**Respostas:**

#### Sucesso (200)
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

#### CEP inválido (422)
```json
{
  "message": "invalid zipcode"
}
```

#### CEP não encontrado (404)
```json
{
  "message": "can not find zipcode"
}
```

### GET /health

Endpoint de health check.

**Exemplo:**
```bash
curl http://localhost:8080/health
```

**Resposta:**
```
OK
```

## 🧪 Testes

Execute os testes automatizados:

```bash
go test -v ./...
```

## 🏗️ Estrutura do Projeto

```
.
├── main.go                      # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/               # Handlers HTTP
│   │   └── weather_handler.go
│   ├── services/               # Serviços de negócio
│   │   ├── cep_service.go      # Integração com viaCEP
│   │   └── weather_service.go  # Integração com WeatherAPI
│   ├── models/                 # Modelos de dados
│   │   └── models.go
│   └── utils/                  # Utilitários
│       └── temperature.go      # Conversões de temperatura
├── Dockerfile                  # Imagem Docker
├── docker-compose.yml          # Configuração Docker Compose
├── go.mod                      # Dependências Go
└── README.md                   # Documentação
```

## 🐳 Docker

### Build da imagem

```bash
docker build -t weather-api .
```

### Executar container

```bash
docker run -p 8080:8080 -e WEATHER_API_KEY=sua-api-key-aqui weather-api
```

## ☁️ Deploy no Google Cloud Run

### Pré-requisitos

1. Instale o [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
2. Configure o projeto:
   ```bash
   gcloud config set project SEU-PROJECT-ID
   ```

### Deploy

1. Build e push da imagem para Google Container Registry:
   ```bash
   # Configurar o Docker para usar gcloud
   gcloud auth configure-docker

   # Build e tag da imagem
   docker build -t gcr.io/SEU-PROJECT-ID/weather-api .

   # Push da imagem
   docker push gcr.io/SEU-PROJECT-ID/weather-api
   ```

2. Deploy no Cloud Run:
   ```bash
   gcloud run deploy weather-api \
     --image gcr.io/SEU-PROJECT-ID/weather-api \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated \
     --set-env-vars WEATHER_API_KEY=sua-api-key-aqui \
     --port 8080
   ```

   Ou usando o arquivo de serviço:
   ```bash
   gcloud run deploy weather-api \
     --source . \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated \
     --set-env-vars WEATHER_API_KEY=sua-api-key-aqui
   ```

3. O Cloud Run retornará a URL do serviço após o deploy.

### Deploy simplificado (a partir do código-fonte)

```bash
gcloud run deploy weather-api \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua-api-key-aqui
```

Este comando fará o build e deploy automaticamente.

## 📚 APIs Utilizadas

- **viaCEP**: https://viacep.com.br/ - Para buscar informações do CEP
- **WeatherAPI**: https://www.weatherapi.com/ - Para buscar temperatura atual

## 🔧 Fórmulas de Conversão

- **Celsius para Fahrenheit**: `F = C * 1.8 + 32`
- **Celsius para Kelvin**: `K = C + 273`

## 📝 Notas

- O CEP pode ser informado com ou sem formatação (ex: `01310100` ou `01310-100`)
- A WeatherAPI oferece um plano gratuito com 1 milhão de requisições por mês
- O Cloud Run Free Tier permite até 2 milhões de requisições por mês

## 📄 Licença

MIT
