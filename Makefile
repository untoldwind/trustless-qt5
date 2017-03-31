PKGS= ./config/... ./ui/... .

VERSION ?= $(shell date -u +%Y%m%d.%H%M%S)VERSION ?= $(shell date -u +%Y%m%d.%H%M%S)

#all: export GOPATH=${PWD}/../../../..
all: format
	@mkdir -p bin
	@echo "--> Running go build ${VERSION}"
	@go build -v -i -o bin/trustless-qt5 github.com/untoldwind/trustless-qt5

minimal: format
	@qtminimal desktop .
	@go build -v -i -tags=minimal -o bin/trustless-qt5 github.com/untoldwind/trustless-qt5


#format: export GOPATH=${PWD}/../../../..
format:
	@echo "--> Running go fmt"
	@go fmt ${PKGS}

glide.install:
	@echo "--> glide install"
	@go get github.com/Masterminds/glide
	@go build -v -o bin/glide github.com/Masterminds/glide
	@bin/glide install -v
