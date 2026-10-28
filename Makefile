.PHONY: build lint test

build:
	go build ./...

lint:
	@files="$$(gofmt -l src tests)"; if [ -n "$$files" ]; then printf 'Files requiring gofmt:\n%s\n' "$$files"; exit 1; fi
	go vet ./...

test:
	go test ./...
