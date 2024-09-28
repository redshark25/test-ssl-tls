FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


# Copy the source code from the main directory
COPY *.go ./

# Build the Go application
RUN go build -o main

# Create the directory for Let's Encrypt certificates
RUN mkdir -p /certs

# Set proper permissions to ensure the app can write to this directory
RUN chmod -R 755 /certs

# Expose the port the app runs on
EXPOSE 80
EXPOSE 443

# Command to run the Go application
CMD ["./main"]