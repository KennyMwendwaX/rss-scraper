build:
	@echo "Building the binary..."
	@go build -o bin/rss ./cmd	

run: build
	@echo "Running the application..."
	@./bin/rss
