# ======================== Etapa 1: Build do Backend ========================
FROM golang:latest AS backend-builder

WORKDIR /app/backend

# cache backend
COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/. .

RUN go get github.com/redis/go-redis/v9

COPY backend/scripts/wait-for-it.sh /app/backend/wait-for-it.sh

# Permissão para execução do script
RUN chmod +x /app/backend/wait-for-it.sh

# Compila back
RUN go build -o /app/backend/fake-fintech ./main.go

# ======================== Etapa 2: Build do Frontend ========================
FROM node:18 AS frontend-builder

WORKDIR /app/frontend

# cache frontend
COPY frontend/package.json frontend/package-lock.json ./

RUN npm install --legacy-peer-deps

COPY frontend/. .

RUN npm run build --prod 

RUN ls -la /app/frontend/dist

# ======================== Etapa 3: Runtime ========================
FROM debian:bullseye-slim

WORKDIR /app

# Instala somente os clientes PostgreSQL e netcat
RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# Copia apenas os arquivos necessários do backend
COPY --from=backend-builder /app/backend/fake-fintech .
COPY --from=backend-builder /app/backend/wait-for-it.sh .

# Copia os arquivos do frontend para servir
COPY --from=frontend-builder /app/frontend/dist/frontend /app/frontend

# Garante que o script tem permissão de execução
RUN chmod +x wait-for-it.sh

EXPOSE 8080

CMD ["./wait-for-it.sh", "postgres:5432", "--", "./wait-for-it.sh", "redis:6379", "--", "./fake-fintech"]
