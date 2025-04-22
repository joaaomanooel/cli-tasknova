# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=tasknova
MAIN_PACKAGE=.

# Output directories
BIN_DIR=./bin
MACOS_BIN=$(BIN_DIR)/$(BINARY_NAME)-darwin
LINUX_BIN=$(BIN_DIR)/$(BINARY_NAME)-linux
WINDOWS_BIN=$(BIN_DIR)/$(BINARY_NAME)-windows.exe

all: test build

build: clean build-macos build-linux build-windows

# Test targets
test:
	gotestsum --format testdox --hide-summary output --format-hide-empty-pkg --packages=all -- -v ./... -coverprofile=coverage/coverage.out -covermode=atomic -coverpkg=./... && \
	go tool cover -html=coverage/coverage.out -o coverage/index.html && \
	go tool cover -func=coverage/coverage.out | grep total: | awk '{if (substr($$3, 1, length($$3)-1) < 90) { print "Coverage " $$3 " is below 90%"; exit 1 } else { print "Coverage " $$3 " meets minimum 90% requirement" }}'

test-coverage:
	go tool cover -func=coverage/coverage.out | grep total: | awk '{printf "Total coverage: %s\n", $$3}'

test-watch: ## Watch for changes and run tests
	gotestsum --watch --format testdox

test-verbose: ## Run tests in verbose mode
	gotestsum --format standard-verbose --packages=all -- -v ./... --cover

test-race: ## Run tests with race detector
	gotestsum --format testdox -- -race ./...

test-bench: ## Run benchmark tests
	gotestsum --format dots -- -bench=. -benchmem ./...

test-nocache: ## Run tests without cache
	gotestsum --format testdox --format-hide-empty-pkg --packages=all -- -count=1 ./...

test-short: ## Run short tests
	gotestsum --format short-verbose -- -short ./...

test-timeout: ## Run tests with timeout
	gotestsum --format testdox -- -timeout 30s ./...

test-failed: ## Watch and run only failed tests
	gotestsum --format testname --packages=all --watch --watch-fail=running -- ./...

clean:
	$(GOCMD) clean
	rm -rf $(BIN_DIR)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

deps:
	$(GOMOD) download

tidy:
	$(GOMOD) tidy

# Create bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Cross compilation targets
build-macos: $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(MACOS_BIN) $(MAIN_PACKAGE)
	@echo "Built for macOS (amd64): $(MACOS_BIN)"

build-macos-arm64: $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)
	@echo "Built for macOS (arm64): $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64"

build-linux: $(BIN_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_BIN) $(MAIN_PACKAGE)
	@echo "Built for Linux: $(LINUX_BIN)"

build-windows: $(BIN_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_BIN) $(MAIN_PACKAGE)
	@echo "Built for Windows: $(WINDOWS_BIN)"

# Create a release with all binaries
release: build
	@echo "Release binaries created in $(BIN_DIR) directory"

.PHONY: all build test test-coverage test-watch test-verbose test-race test-bench test-nocache test-short test-timeout test-failed clean run deps tidy build-macos build-macos-arm64 build-linux build-windows release