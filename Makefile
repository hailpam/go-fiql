
all: clean build test build-examples

clean:
	rm -f examples/ast

build:
	go build ./gofiql

test:
	go test ./gofiql

build-examples:
	go build -o examples/ast examples/ast.go

run-examples:
	go run examples/ast.go

.PHONY: build test all run-examples build-examples
