BINARY_NAME=main
 
all: build test
 
build:
	go build -o .out/${BINARY_NAME} cmd/main.go
 
test:
	go test -v cmd/main.go
 
run:
	go build -o .out/${BINARY_NAME} cmd/main.go
 
deps:
	go get github.com/gin-gonic/gin@v1.10.0
	go get github.com/go-sql-driver/mysql@v1.8.1
	go get github.com/jmoiron/sqlx@v1.4.0

dev:
	air

clean:
	go clean
	rm ${BINARY_NAME}
