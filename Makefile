clean:
	rm -Rf parsing/ bin/

test:
	go test ./...

gen-parser: clean
	go generate ./...

build: clean gen-parser
	go build -o bin/ .

compile-and-run-example:
	rm -f $(program).jasm $(program).class
	./bin/poj ./examples/$(program)
	jasm $(program).jasm
	java $(program)
