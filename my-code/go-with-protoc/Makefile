generate:
	./generate.sh

run:
	go fmt
	go run ./...

build:
	rm -rf ./build
	go fmt ./main.go
	go build ./...
	go build -o ./build/go-with-protoc main.go

run-build:
	./build/go-with-protoc

execute: build run-build