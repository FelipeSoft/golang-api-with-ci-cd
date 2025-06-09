FROM golang:current

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o golang-api-with-ci-cd ./cmd/golang-api-with-ci-cd

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/golang-api-with-ci-cd .

CMD ["./golang-api-with-ci-cd"]
