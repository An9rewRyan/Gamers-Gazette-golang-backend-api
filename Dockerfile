# Build the Go API
FROM golang:latest AS builder
WORKDIR "/backend"
RUN ls
ADD backend .
RUN go mod download
RUN go mod tidy
RUN go build .
CMD ./go
