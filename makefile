.PHONY: build run

# Build the project and specify the output directory and file name
build:
	go build -o ./bin/myproject

# Build the project and run the binary
run: build
	./bin/myproject
