# Build the Go API
FROM golang:latest AS builder
WORKDIR "/backend"
RUN ls
ADD backend .
RUN go mod download
RUN go mod tidy
RUN go build .
CMD ./go

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# Build the React application
FROM node:latest AS node_builder
WORKDIR "/frontend"
RUN ls
ADD frontend .
RUN ls
RUN npm install
RUN npm run build
RUN serve -s build
# npm install -g serve

# RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM ubuntu:latest
WORKDIR "/main"
# RUN apk --no-cache add ca-certificates
COPY --from=builder "/backend" ./backend
COPY --from=node_builder "/frontend" ./frontend
# RUN apt-get update
# RUN apt-get install -y supervisor
# ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf 
EXPOSE 8080
CMD cd backend; ./go