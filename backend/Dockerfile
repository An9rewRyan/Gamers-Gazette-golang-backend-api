FROM golang:latest AS builder
WORKDIR "/backend"
RUN ls
ADD backend .
RUN ls
RUN go mod download
RUN go mod tidy
RUN go build .
EXPOSE 8000
CMD ./go