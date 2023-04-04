# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

WORKDIR /go/src/golangdebug

COPY . .

RUN go mod download && \
    go mod tidy

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build main.go

EXPOSE 2345
EXPOSE 8080

CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "./main", "--check-go-version=false"]
