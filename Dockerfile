FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git curl bash

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 3000

CMD ["air"]