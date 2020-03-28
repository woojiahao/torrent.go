FROM golang:1.14.0-buster
WORKDIR /app
COPY . .
RUN go mod download
RUN go test test/*/*.go
