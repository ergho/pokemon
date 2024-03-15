build:
	@go build -o bin/pokedex

run: build
	@./bin/pokedex

test:
	@go test -v ./...
