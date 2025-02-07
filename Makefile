# (CGO_ENABLED=0 because my application doesn't use C libraries)

BINARY_NAME=main

all: build test
 
build:
	CGO_ENABLED=0	go build -o .out/${BINARY_NAME} cmd/main.go

test:
	go test -v cmd/main.go
 
run:
	CGO_ENABLED=0	go build -o .out/${BINARY_NAME} cmd/main.go
	./.out/${BINARY_NAME}
 
deps:
	go mod download	

dev:
	air

clean:
	go clean
	rm ${BINARY_NAME}
