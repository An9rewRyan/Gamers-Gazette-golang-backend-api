# FROM node:latest AS node_builder
# WORKDIR "/frontend"
# RUN ls
# ADD frontend .
# RUN ls
# RUN npm install
# RUN npm run build
# RUN ls
# npm install -g serve

# Build the Go API
FROM golang:latest AS builder
WORKDIR "/backend"
RUN ls
ADD backend .
RUN go mod download
RUN go mod tidy
RUN go build .

# RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM ubuntu:latest
WORKDIR "/main"
# RUN apk --no-cache add ca-certificates
COPY --from=builder "/backend" ./backend
RUN ls
# COPY --from=node_builder "frontend/build" ./web
# RUN ls
# RUN apt-get update
# RUN apt-get install -y supervisor
# ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf 
EXPOSE 3000
CMD cd backend; ./go