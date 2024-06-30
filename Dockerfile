FROM golang:1.22.2 AS builder

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go build -o /app/bin/web /app/cmd/web/main.go

EXPOSE 1337 

CMD ["air", "-c", ".air.toml"]
