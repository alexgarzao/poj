clean:
	rm -Rf parsing/ bin/

test:
	go clean -testcache
	go test -v ./...

gen-parser: clean
	go generate ./...

build: clean gen-parser
	go build -o bin/ .

compile-and-run-example:
	rm -f $(program).jasm $(program).class
	./bin/poj ./tests/valid_pascal_programs/$(program)
	jasm $(program).jasm
	java $(program)
