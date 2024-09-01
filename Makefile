# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o tmp/main cmd/api/main.go
	@swag init -g ./cmd/api/main.go -d ./,./internal/models

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
# test:
# 	@echo "Testing..."
# 	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f tmp/main

watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching..."; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
			echo "Watching..."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi
