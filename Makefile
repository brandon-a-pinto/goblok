build:
	@go build -o bin/goblok -C cmd/blockchain

run: build
	@./cmd/blockchain/bin/goblok
