FROM golang:1.23.6-alpine3.21 AS builder

WORKDIR /globetrotter

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest

WORKDIR /globetrotter

# Copy the binary and config file from the builder stage
COPY --from=builder /globetrotter/server .
COPY --from=builder /globetrotter/app.yaml ./app.yaml

# Expose application port (change if necessary)
EXPOSE 8080

# Run the binary
CMD ["./server"]
ARG ["--config","./app.yaml"]