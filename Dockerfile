# Etapa 1: Build
FROM golang

WORKDIR /app

# Copia apenas arquivos essenciais para maximizar cache
COPY go.mod go.sum ./
RUN go mod download

# Copia o código restante
COPY . .

# Compila a aplicação
RUN go build -o reusable-api

# Etapa 2: Runtime
FROM debian:bullseye-slim

WORKDIR /app

# Instala somente os clientes PostgreSQL e netcat
RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd \
    && rm -rf /var/lib/apt/lists/*

# Copia apenas os arquivos necessários da build anterior
COPY --from=builder /app/reusable-api .
COPY --from=builder /app/scripts/wait-for-it.sh .

# Garante que o script tem permissão de execução
RUN chmod +x wait-for-it.sh

EXPOSE 8080

CMD ["./wait-for-it.sh", "postgres:5432", "--", "./wait-for-it.sh", "redis:6379", "--", "./reusable-api"]
