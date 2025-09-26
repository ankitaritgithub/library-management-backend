# Use the official Golang image
FROM golang:1.22-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Enable CGO
ENV CGO_ENABLED=1

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with CGO enabled
RUN CGO_ENABLED=1 go build -o main .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
