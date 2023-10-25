# STEP 1: build executable binary

# Use the official Golang image as the build stage
FROM golang:1.18-alpine AS build

# Set a label for the maintainer
LABEL maintainer="Titanio Yudista <titanioyudista98@gmail.com>"

# Set the working directory in the build stage
WORKDIR /app

# Install required dependencies for building the Go application
RUN apk add --no-cache bash make gcc libc-dev

# Copy Go module files and build cache for better Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and any additional files
COPY . .

# Copy the environment file (you should provide this file)
RUN cp .env.example .env

# Build the Go application
RUN go build -o cakeplabs-exam

# Create a vendor directory and copy dependencies
RUN go mod vendor

# STEP 2: create a smaller image for the final application

# Use a minimal Alpine Linux image as the final stage
FROM alpine:3.16.0

# Install necessary packages for the final image
RUN apk add --no-cache bash

# Copy the compiled binary and vendor directory from the build stage
COPY --from=build /app/cakeplabs-exam /
COPY --from=build /app/vendor /vendor

# Expose port if your application listens on a specific port
EXPOSE 8000

# Define the command to run your application
CMD ["/cakeplabs-exam"]
