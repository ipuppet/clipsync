.PHONY: default build run clean help

BINARY="clipsync.exe"

default: build

export CGO_ENABLED=0
export GOARCH=amd64

build: export GOOS=windows
build: clean
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	go generate
	go build -ldflags="-s -w -H=windowsgui" -o ${BINARY}

run: export GOOS=windows
run:
	go run .

clean:
	go clean
	-@rm -rf ./${BINARY}

help:
	@echo "make           - Default build"
	@echo "make build     - Build binary for Windows"
	@echo "make run       - Run code use `go run`"
	@echo "make clean     - Clean