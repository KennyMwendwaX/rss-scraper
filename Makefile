build:
	@go build -o bin/rss	

run: build
	@./bin/rss
