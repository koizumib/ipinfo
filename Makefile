APP_NAME := ipinfo
PKG := ./...
CMD := ./cmd/$(APP_NAME)

.PHONY: tidy build run test lint air

tidy:
	go mod tidy

build:
	go build -o bin/$(APP_NAME) $(CMD)

run:
	go run $(CMD) 192.168.100.100/24

test:
	go test -v $(PKG)

lint:
	golangci-lint run
