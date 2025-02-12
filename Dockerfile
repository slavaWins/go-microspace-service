FROM golang:1.23-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
#RUN go build -o main .

RUN CGO_ENABLED=0 go build -o .


EXPOSE 8081



CMD ["go", "run", "main.go"]