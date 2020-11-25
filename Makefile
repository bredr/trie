all: install-tools lint test

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.32.2

lint:
	@echo golangci-lint run
	@./bin/golangci-lint run

test:
	@echo go test ./...
	@go test ./...

bench:
	@echo Running benchmarks...
	@go test -run=. -bench=. -benchtime=5s -count 2 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out | tee bench.txt
	@echo For further analysis...
	@echo go tool pprof -http :8080 cpu.out
	@echo go tool pprof -http :8081 mem.out
	@echo go tool trace trace.out
	@echo go run golang.org/x/perf/cmd/benchstat bench.txt