FROM golang:latest AS builder
WORKDIR "/backend"
ADD backend .
RUN ls
RUN go mod download
RUN go mod tidy
RUN go build .
EXPOSE 8000
CMD ./go