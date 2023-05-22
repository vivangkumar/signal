.PHONY: fmt lint test

lint:
	go vet  ./...
	golint ./...

fmt:
	go fmt -s -w .

test:
	go test ./... -v
