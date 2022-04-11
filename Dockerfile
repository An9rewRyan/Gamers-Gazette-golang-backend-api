# Build the Go API
FROM golang:latest AS builder
WORKDIR "/backend"
RUN ls
ADD backend .
RUN ls
# RUN rm -r go
RUN go mod download
RUN go mod tidy
RUN go build .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# Build the React application
FROM node:alpine AS node_builder
WORKDIR "/frontend"
RUN ls
ADD frontend .
RUN ls
RUN npm install
# RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
WORKDIR /main
RUN apk --no-cache add ca-certificates
COPY --from=builder . .
COPY --from=node_builder . .
RUN ls
# RUN chmod +x ./main
EXPOSE 8000
EXPOSE 3000
RUN ls
CMD cd backend; ./go