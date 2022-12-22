VERSION := latest

build:
	docker build -t ton-http-liteapi:$(VERSION) --target ton-http-liteapi .