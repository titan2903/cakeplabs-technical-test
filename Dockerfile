# STEP 1: build executable binary

# Use the official Golang image as the build stage
FROM golang:alpine3.16 AS build

# Install git, required for fetching Go dependencies.
# Utilize multi-stage builds to keep the final image clean and small.
RUN apk update && apk add --no-cache git

# Set a label for the maintainer
LABEL maintainer="Titanio Yudista <titanioyudista98@gmail.com>"

# Set the working directory in the build stage
WORKDIR /app

# Copy the go module files and download dependencies for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and any additional files
COPY . .

# Copy .env
COPY .env.example .env

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cakeplabs-technical-test .

# STEP 2: create a smaller image for the final application
# Use a minimal Alpine Linux image as the final stage
FROM alpine:latest

# Add ca-certificates in case your application makes external HTTPS calls
# Ensure the application runs securely and efficiently
RUN apk --no-cache add ca-certificates

# Copy the compiled application from the build stage to the final image
COPY --from=build /app/cakeplabs-technical-test .
COPY --from=build /app/.env .

# Expose port if your application listens on a specific port
EXPOSE 8090

# Define the command to run your application
CMD ["./cakeplabs-technical-test"]