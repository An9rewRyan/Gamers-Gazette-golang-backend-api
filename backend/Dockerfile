FROM golang:latest AS builder
RUN ls
WORKDIR /backend
RUN ls
ADD . .
RUN ls
RUN go mod download
RUN go mod tidy
RUN go build .
CMD ./go