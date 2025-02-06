# Use the latest Go image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN make deps

# Copy the source code from the host to the container
COPY . .

# Build the application inside the container
RUN make build && chmod +x .out/main

EXPOSE 8080

# Set the command to run the application
CMD ["/app/.out/main"]

