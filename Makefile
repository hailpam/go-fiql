
all: build test

build:
	go build .

test:
	go test .

.PHONY: build test all
