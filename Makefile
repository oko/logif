GO111MODULES := on

all: test coverage lint fmt vet

.PHONY: test
test:
	go test

.PHONY: coverage
coverage:
	go test -coverprofile=/tmp/coverage.out
	go tool cover -func=/tmp/coverage.out

.PHONY: lint
lint:
	find . -type f -name '*.go' | xargs golint

.PHONY: fmt
fmt:
	find . -type f -name '*.go' | xargs gofmt -w -s

.PHONY: vet
vet:
	go vet
