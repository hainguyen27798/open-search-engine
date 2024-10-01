run:
	air

wire:
	cd internal/wires && wire

build:
	go build -o ./bin/main ./server

test:
	go clean -testcache && go test ./tests -v
