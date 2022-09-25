run:
	go run -race .

test:
	go test -v -race -cover ./...