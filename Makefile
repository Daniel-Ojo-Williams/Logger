build:
	@go build -o bin/ .

run:
	@go run .

test:
	@go test ./...
