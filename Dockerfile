# Start from the official Go image
FROM golang:1.23.0

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Get the swag tool
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag


# Build the application
# RUN go build -o main .
RUN make build-prod

# Expose port 8080 for the Gin server
EXPOSE 8080

# Run the application
SHELL ["/bin/bash", "-c"]
# ENTRYPOINT make run
CMD ["./tmp/main"]
