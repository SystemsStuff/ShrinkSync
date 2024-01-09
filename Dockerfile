FROM golang:1.21-alpine

WORKDIR /src/shrinksync

COPY . .

EXPOSE 8080

RUN apk update && apk add curl jq

RUN go build -o shrinksync sampleServer/server/main.go

CMD ["./shrinksync"]
