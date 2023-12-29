# Stage 1: Build the Go binary
FROM golang:1.17 AS builder

WORKDIR /app

# Copy only the Go mod and sum files to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o myapp .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/myapp .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the executable
CMD ["./myapp"]
