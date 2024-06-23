# Makefile

BUILD_DIR=build
BINARY_NAME=stump
LINUX_BINARY=$(BUILD_DIR)/$(BINARY_NAME)
WINDOWS_BINARY=$(BUILD_DIR)/$(BINARY_NAME).exe

build: build-linux build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY) ./cmd

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BINARY) ./cmd

run-linux:
	./build/stump

run-windows:
	./build/stump.exe
