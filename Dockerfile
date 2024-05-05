# Use the official Go image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Install CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon

# Set the entry point for the container
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o app" -command="./app"