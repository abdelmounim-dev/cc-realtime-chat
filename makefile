.PHONY: all clean build

# Directories
BINDIR := bin

# Binaries
SERVER_BINARY := $(BINDIR)/server
CLIENT_BINARY := $(BINDIR)/client

all: clean build

# Build both server and client
build: $(SERVER_BINARY) $(CLIENT_BINARY)

# Build the server binary
$(SERVER_BINARY): cmd/server/main.go
	@echo "Building server..."
	go build -o $(SERVER_BINARY) cmd/server/main.go

# Build the client binary
$(CLIENT_BINARY): cmd/client/main.go
	@echo "Building client..."
	go build -o $(CLIENT_BINARY) cmd/client/main.go

# Clean up the binaries
clean:
	@echo "Cleaning up..."
	rm -f $(SERVER_BINARY) $(CLIENT_BINARY)
