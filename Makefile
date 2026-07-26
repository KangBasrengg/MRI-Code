APP_NAME := codemri
BUILD_DIR := bin
ENTRY := ./cmd/codemri

.PHONY: all build clean test run doctor version

all: clean build

build:
	@echo "🔨 Compiling CodeMRI v0.1.0 Genesis..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY)
	@echo "✨ Build completed successfully! Binary available at: $(BUILD_DIR)/$(APP_NAME)"

clean:
	@echo "🧹 Cleaning build output and caches..."
	@rm -rf $(BUILD_DIR)
	@go clean

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

version: build
	@./$(BUILD_DIR)/$(APP_NAME) version

doctor: build
	@./$(BUILD_DIR)/$(APP_NAME) doctor
