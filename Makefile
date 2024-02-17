clean:
	rm -Rf parsing/ bin/

test:
	go test ./...

gen-parser: clean
	go generate ./...

build: clean gen-parser
	go build -o bin/ .
