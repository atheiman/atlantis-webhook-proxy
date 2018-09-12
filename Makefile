ZIP_TARGET := $(shell basename ${PWD}.zip)

dep:
	dep ensure -v

build:
	env GOOS=linux go build -v -ldflags="-s -w" -o bin/api api/main.go

test:
	go test -v ./...

package:
	zip $(ZIP_TARGET) bin/*
