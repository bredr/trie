all: install-tools generate lint test

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.32.2

lint:
	@echo golangci-lint run
	@./bin/golangci-lint run

generate:
	@echo go generate ./...
	@go generate ./...

test:
	@echo go test ./...
	@go test ./...
	@go test -bench=.