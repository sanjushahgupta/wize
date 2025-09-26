# Use an official Go image (includes Go toolchain + runtime)
FROM golang:1.24.4

# Create and switch to a working directory inside the image
WORKDIR /app

# First copy go.mod/sum and download deps (caches better)
COPY go.mod go.sum ./
RUN go mod download

# Then copy the rest of your source code
COPY . .

# Build your app (outputs a single binary named 'app')
RUN go build -o app ./cmd/main.go

# The port your API listens on inside the container
EXPOSE 8080

# Run the binary
CMD ["./app"]
