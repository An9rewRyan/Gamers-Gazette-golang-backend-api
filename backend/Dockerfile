FROM golang:1.17-alpine

WORKDIR /backend

COPY . .
RUN go mod download
RUN go mod tidy
