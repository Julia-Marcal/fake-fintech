FROM golang:latest AS backend-builder

WORKDIR /app/backend

# cache backend
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd
RUN go get github.com/redis/go-redis/v9

COPY scripts/wait-for-it.sh /app/backend/wait-for-it.sh

# Permissão para execução do script
RUN chmod +x /app/backend/wait-for-it.sh

RUN go build

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD [ "./backend/wait-for-it.sh", "postgres:5432", "--", "./backend/wait-for-it.sh", "redis:6379", "--", "./fake-fintech"]