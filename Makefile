.PHONY: build-gofumpt
build-gofumpt:
	go build -modfile=tools/go.mod -o bin/gofumpt mvdan.cc/gofumpt

.PHONY: build-golangci
build-golangci:
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: check
check:
	if [ ! -f "bin/golangci-lint" ]; then make build-golangci; fi
	bin/golangci-lint run

FILES = $(shell find . -type f -name '*.go')

.PHONY: format
format:
	if [ ! -f "bin/gofumpt" ]; then make build-gofumpt; fi
	bin/gofumpt -w $(FILES)

.PHONY: test
test:
	go test -cover -coverprofile=coverage.out ./core/...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: all
all:
	go mod tidy
	make check
	make format
	make test
