build:
	go install ./...

install:
	go install ./...

test:
	go test ./...

lint:
	golangci-lint run

generate:
	go generate

check-clean:
	@git diff --exit-code || (echo "\033[0;31mWorking directory is not clean - did you run 'make generate' and commit the changes?" && exit 1)

.PHONY: install build test lint generate check-clean
