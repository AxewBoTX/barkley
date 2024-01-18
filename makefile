GO := go
BINARY_NAME := dondu
BUILD_DIR := build

build:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .
	$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)

.PHONY: all build clean
