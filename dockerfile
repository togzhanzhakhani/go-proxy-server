# Use the official Golang image to create a build artifact.
# This container will include the source code and compile the application.
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o go-proxy-server .

# Start a new stage from scratch
FROM debian:bullseye-slim

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/go-proxy-server /go-proxy-server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go-proxy-server"]
