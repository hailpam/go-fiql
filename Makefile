
travis: clean build test

all: clean build test build-examples

clean:
	rm -f examples/ast/ast
	rm -f examples/visit/visit

build:
	go build ./gofiql

test:
	go test ./gofiql

build-examples:
	go build -o examples/ast/ast examples/ast/ast.go
	go build -o examples/visit/visit examples/visit/visit.go

run-examples:
	go run examples/ast.go

.PHONY: travis build test all run-examples build-examples
