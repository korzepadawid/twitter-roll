run:
	go run -race .

test:
	go test -v -race -cover ./...

build:
	rm -rf backend.out && go build -o ./backend.out