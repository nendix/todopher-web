BINARY_NAME=main
 
all: build test
 
build:
	go build -o .out/${BINARY_NAME} cmd/main.go
 
test:
	go test -v cmd/main.go
 
run:
	go build -o .out/${BINARY_NAME} cmd/main.go
	./.out/${BINARY_NAME}
 
deps:
	go mod download	

dev:
	air

clean:
	go clean
	rm ${BINARY_NAME}
