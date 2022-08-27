.PHONY: run
run:
	go run ./cmd/bot/main.go

.PHONY: build
build:
	go build -o bot-runner ./cmd/bot/main.go