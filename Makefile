clean:
	rm -Rf parsing/ bin/

gen-parser: clean
	go generate ./...

build: clean gen-parser
	go build -o bin/ .
