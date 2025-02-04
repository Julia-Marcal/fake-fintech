# ======================== Etapa 1: Build do Backend ========================
FROM golang:latest AS backend-builder

WORKDIR /app/backend

# Copia apenas arquivos essenciais para maximizar cache
COPY backend/go.mod backend/go.sum

# Debug: Check if the files exist after copying
RUN ls -la /app/backend

# Baixa as dependências do Go
RUN go mod download

# Copia o restante do código
COPY backend/. .

# Instala dependências do PostgreSQL e netcat
RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# Instala dependências do Redis
RUN go get github.com/redis/go-redis/v9

# Copia o script wait-for-it.sh
COPY backend/scripts/wait-for-it.sh /app/backend/wait-for-it.sh

# Permissão para execução do script
RUN chmod +x /app/backend/wait-for-it.sh

# Compila a aplicação
RUN go build -o fake-fintech

# ======================== Etapa 2: Build do Frontend ========================
FROM node:18 AS frontend-builder

WORKDIR /app/frontend

# Copia apenas arquivos essenciais para maximizar cache
COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

# Copia o restante do código
COPY frontend/. .

# Compila a aplicação Angular
RUN npm run build --prod

# ======================== Etapa 3: Runtime ========================
FROM debian:bullseye-slim

WORKDIR /app

# Instala somente os clientes PostgreSQL e netcat
RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# Copia apenas os arquivos necessários do backend
COPY --from=backend-builder /app/backend/fake-fintech .
COPY --from=backend-builder /app/wait-for-it.sh .

# Copia os arquivos do frontend para servir
COPY --from=frontend-builder /app/frontend/dist/frontend /app/frontend

# Garante que o script tem permissão de execução
RUN chmod +x wait-for-it.sh

EXPOSE 8080

CMD ["./wait-for-it.sh", "postgres:5432", "--", "./wait-for-it.sh", "redis:6379", "--", "./fake-fintech"]