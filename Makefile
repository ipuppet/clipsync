.PHONY: default build build_win run gotool clean help

BINARY="clipsync"
BINARY_WIN="clipsync.exe"

default: gotool build

export CGO_ENABLED=0
export GOARCH=amd64

build: export GOOS=linux
build: clean
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	go build -ldflags="-s -w" -o ${BINARY}

build_win: export GOOS=windows
build_win: clean
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	go build -ldflags="-s -w" -o ${BINARY_WIN}

run: export GOOS=windows
run: export GO_HOST=localhost
run:
	@go env -w CGO_ENABLED=$(CGO_ENABLED)
	@go env -w GOOS=$(GOOS)
	@go env -w GOARCH=$(GOARCH)
	@go run ./ -log=./var

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make build_win - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make clean - 移除二进制文件和 vim swap files"