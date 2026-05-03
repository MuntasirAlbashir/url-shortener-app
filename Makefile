.PHONY: build test clean

build:
	go build -o ./bin/api ./cmd/api/

test:
	go test -v -race ./...

clean:
	rm -rf ./bin/

