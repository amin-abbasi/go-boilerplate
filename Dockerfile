FROM golang:1.20-alpine

WORKDIR /opt/go_boilerplate

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@v1.27.8

ENTRYPOINT air
