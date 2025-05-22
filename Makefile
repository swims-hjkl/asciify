BIN_DIR := bin
CLI_NAME := asciifycli
CLI_PATH := ./cmd/asciifycli

build:
	mkdir -p $(BIN_DIR)
	go test
	go build -o $(BIN_DIR)/$(CLI_NAME) $(CLI_PATH)

benchmark:
	go test -bench=.
