FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o golang-api .

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/golang-api .

CMD ["./golang-api"]
