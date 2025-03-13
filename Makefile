# Format code 
fmt:
	go fmt ./...

# View possible issues in codebase
vet:
	go vet ./...

# Add any missing libraries and remove unsed ones
tidy: fmt
	go mod tidy

# Start the go server
start:
	go run ./cmd

# View the makefile commads
view:
	@cat Makefile