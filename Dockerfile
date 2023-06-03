# syntax=docker/dockerfile:1

# Intial build stage
FROM golang:1.20 AS build-stage

# Set destination for COPY
WORKDIR /app

# Copy dependency files and download
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy all excpet dockerignore and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd/users/ -o users 

# Deploy the application binary into a lean image
# FROM alpine:latest AS build-release-stage
# WORKDIR /
# COPY --from=build-stage app/cmd/users/users /users

EXPOSE 8080

# Run
# CMD ["/users"]
CMD ["./cmd/users/users"]
