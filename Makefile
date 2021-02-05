# .PHONY 有 build 文件，不影响 build 命令执行
.PHONY: build install build-linux run clean docs
# 二进制文件名
BUILDDIR="build"
# 二进制文件名
PROJECT="smileHugo"
# 启动文件PATH
MAIN_PATH="main.go"
VERSION="0.0.1"
DATE= `date +%FT%T%z`

version:
	@echo version: ${VERSION}

install:
	@echo download package
	@go mod download

# 交叉编译运行在linux系统环境
build-linux:
	@echo version: ${VERSION} date: ${DATE} os: linux
	@GOOS=linux go build -o ${BUILDDIR}/${PROJECT} ${MAIN_PATH}

run: build
	@./${BUILDDIR}/${PROJECT}

clean:
	rm -rf ./build
	rm -rf ./logs