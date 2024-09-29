run:
	air

wire:
	cd internal/wires && wire

build:
	go build -o ./bin/main ./server
