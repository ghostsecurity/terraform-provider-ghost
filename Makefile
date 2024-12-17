build:
	go install ./...

install:
	go install ./...

test:
	go test ./...

test-acc:
	TF_ACC=true go test ./...

lint:
	golangci-lint run

openapi-client:
	which openapi-generator-cli 2>/dev/null || npm install -g @openapitools/openapi-generator-cli 
	openapi-generator-cli version-manager set 7.10.0
	openapi-generator-cli generate -c openapi-client-config.json

generate: openapi-client
	go generate

check-clean:
	@git diff --exit-code || (echo "\033[0;31mWorking directory is not clean - did you run 'make generate' and commit the changes?" && exit 1)

.PHONY: install build test lint generate check-clean test-acc openapi-client
