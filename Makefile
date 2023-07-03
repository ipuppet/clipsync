.PHONY: default build build_win run clean help

BINARY="clipsync"
BINARY_WIN="clipsync.exe"

default: build_win

export CGO_ENABLED=0
export GOARCH=amd64

build: export GOOS=linux
build: clean
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	go generate
	go build -ldflags="-s -w" -o ${BINARY}

build_win: export GOOS=windows
build_win: clean
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	go generate
	go build -ldflags="-s -w -H=windowsgui" -o ${BINARY_WIN}

run: export GOOS=windows
run:
	go run ./ -addr=localhost

clean:
	go clean

help:
	@echo "make           - Default build (build_win)"
	@echo "make build     - Build binary for linux"
	@echo "make build_win - Build binary for Windows"
	@echo "make run       - Run code use `go run`"
	@echo "make clean     - Run `go clean`