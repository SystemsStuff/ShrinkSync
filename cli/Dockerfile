# syntax=docker/dockerfile:1
FROM golang:1.21-alpine

WORKDIR /src/shrinksync

COPY . .

EXPOSE 8080

RUN go build -o shrinksync main.go

CMD ["./shrinksync"]