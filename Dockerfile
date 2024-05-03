# Use an official Golang runtime as the base image
FROM golang:1.21.6 as builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any needed packages specified in go.mod
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN ls /app

RUN chmod +x main
# Start a new stage from scratch
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the main executable from the builder stage
COPY --from=builder /app/main .

# Install any needed packages
RUN apk add --no-cache redis ca-certificates

# Expose the port that the app listens on
EXPOSE 8080 6379

# Set the environment variables
ENV PORT=:8080 \
  REDIS_ENDPOINT=localhost:6379 \
  REDIS_DB= \
  REDIS_PASSWORD=0 \
  SHORT_URL_DOMAIN=http://localhost:8080/

# Run the app when the container launches
CMD redis-server --daemonize yes && ./main
