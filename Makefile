tidy:
	go mod tidy

test:
	go test ./... -v -cover

lint:
	golangci-lint --version
	golangci-lint run

all: tidy test lint
