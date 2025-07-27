.PHONY: build run test clean install deps migrate

# Build the application
build:
	go build -o bin/script-center main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f script-center.db

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run database migrations
migrate:
	go run main.go migrate

# Install the application
install: build
	sudo cp bin/script-center /usr/local/bin/

# Development setup
dev-setup: deps
	@echo "Setting up development environment..."
	@if [ ! -f .script-center.yaml ]; then \
		echo "Creating default config file..."; \
		cp .script-center.yaml.example .script-center.yaml 2>/dev/null || true; \
	fi

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate documentation
docs:
	@echo "Generating documentation..."
	go run main.go --help > docs/cli-help.txt

# Docker build
docker-build:
	docker build -t script-center .

# Docker run
docker-run:
	docker run -p 8080:8080 -v $(PWD)/data:/app/data script-center

# Quick setup for development
setup: deps migrate create-admin
	@echo "Development environment setup complete!"

# Create default admin user
create-admin:
	@echo "Creating default admin user..."
	@./bin/script-center user create --username admin --role admin --password admin123 || true

# Create example scripts
create-examples: build
	@echo "Creating example scripts..."
	@./bin/script-center script create --name hello-bash --language bash --description "Hello world bash script" --tags example,bash || true
	@./bin/script-center script create --name hello-python --language python --description "Hello world python script" --tags example,python || true

# Run all tests including integration tests
test-all: test
	@echo "Running integration tests..."
	@go test -v ./test/...

# Check for security vulnerabilities
security-check:
	@echo "Running security checks..."
	@go list -json -m all | nancy sleuth

# Performance benchmark
benchmark:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Show help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run unit tests"
	@echo "  test-all       - Run all tests including integration"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  migrate        - Run database migrations"
	@echo "  setup          - Quick development setup"
	@echo "  create-admin   - Create default admin user"
	@echo "  create-examples- Create example scripts"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  docs           - Generate documentation"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  help           - Show this help"