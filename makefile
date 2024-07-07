GOROUTINES = 1
FILE = data.json

.PHONY: all build run test clean generate

all: build

build:
	@echo "Building the application..."
	@go build -o main ./cmd/main.go

run: build
	@echo "Running the application..."
	@./main -goroutines=$(GOROUTINES) -file=$(FILE)

test:
	@echo "Running tests..."
	@go test ./...

generate:
	@echo "Generating test data..."
	@go run tools/datagen/main.go -numObjects=1000000 -file=$(FILE)