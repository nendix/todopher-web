# Use the latest Go image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code from the host to the container
COPY . .

# Build the application inside the container
RUN go build -o .out/main cmd/main.go

# Set executable permissions for the binary (optional if building inside the container)
RUN chmod +x .out/main

# Set the command to run the application
CMD ["/app/.out/main"]
