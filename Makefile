gen-parser:
	rm -Rf parsing/ && go generate ./...

run:
	go run main.go
