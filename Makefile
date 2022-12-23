VERSION := latest

build:
	docker build -t kosrk/ton-http-liteapi:$(VERSION) --target ton-http-liteapi .