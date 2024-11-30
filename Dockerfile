# Use a minimal base image to build the Go application
FROM golang:1.23-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application inside the container
RUN go build -o receipt-processor .

# Use a minimal image to run the application
FROM debian:bullseye-slim

# Install dependencies required for running the app (e.g., glibc)
RUN apt-get update && apt-get install -y libc6

# Set the working directory inside the final container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/receipt-processor .

# Expose the port that the application will listen on
EXPOSE 8080

# Command to run the application
CMD ["./receipt-processor"]
