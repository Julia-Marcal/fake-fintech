FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o fake-fintech

EXPOSE 8080

CMD ["./fake-fintech"]