.PHONY: lint install-linter test format setup prep

lint: install-linter
	golangci-lint run -v ./...

install-linter:
	which golangci-lint || go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

test:
	go test ./...

format:
	go fmt ./...

setup: install-linter
	go mod download

prep: format lint test
	rm -rf ./bin
	go mod tidy
