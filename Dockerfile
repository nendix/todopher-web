# Use the official Golang image to build the Go app
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o todo-app main.go

# Use a smaller image to run the Go app
FROM alpine:latest
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/todo-app .

# Expose the application on port 8080
EXPOSE 8080

# Command to run the app
CMD ["./todo-app"]
