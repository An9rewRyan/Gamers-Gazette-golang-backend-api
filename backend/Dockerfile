FROM golang:1.17-alpine

WORKDIR /backend

COPY . .
RUN rm -r go
RUN go mod download
RUN go mod tidy
