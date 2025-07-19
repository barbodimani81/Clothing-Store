# syntax=docker/dockerfile:1.4
FROM golang:1.24

# Set the working directory
WORKDIR /app

# Copy go mod and sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./main"]
