GO111MODULES := on

test:
	go test

coverage:
	go test -coverprofile=/tmp/coverage.out
	go tool cover -func=/tmp/coverage.out
