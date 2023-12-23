# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = main

# Default target: build the executable
all: $(TARGET)

build:
	go build main.go

# Rule to build the target executable and run it
run: $(TARGET)
	templ generate --watch
	air $(TARGET)
# go build main.go
# ./$(TARGET)

generate:
	templ generate

# Clean target: remove the target executable
clean:
	rm -f $(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...