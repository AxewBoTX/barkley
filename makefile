GO := go
BINARY_NAME := dondu
BUILD_DIR := build

all: build

build:
	@make clean
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@rm -rf $(BUILD_DIR)

.PHONY: all build clean
