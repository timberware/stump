# Makefile

BINARY_NAME=stump
BUILD_DIR=build
LINUX_BINARY=$(BUILD_DIR)/$(BINARY_NAME)
WINDOWS_BINARY=$(BUILD_DIR)/$(BINARY_NAME).exe

build: build-linux build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY) ./cmd/stump

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BINARY) ./cmd/stump
