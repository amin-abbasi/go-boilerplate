FROM golang:1.18-alpine

WORKDIR /opt/go-app
COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT air
