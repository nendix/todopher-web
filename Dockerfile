# Stage 1: Build
FROM golang:latest AS builder

# Install make in the build environment
RUN apt-get update && apt-get install -y make

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN make deps

# Build the application
RUN make build

# Stage 2: Run
FROM scratch

# Set working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=builder /app/.out/main .

# Copy template and static files
COPY --from=builder /app/web/ /app/web

# Expose the application port(Documentation purposes)
EXPOSE 8080

# Run the application
CMD ["./main"]
