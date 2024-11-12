# Build env & run
FROM golang:1.23-alpine

# Set work dir app
WORKDIR /app
# Copy current all file to app
COPY . .
# Go build App
RUN go build ./cmd/main/main.go
# Run app
CMD ["./main"]