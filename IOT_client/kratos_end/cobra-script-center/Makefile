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