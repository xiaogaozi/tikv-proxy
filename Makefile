all: gen-proto server example

gen-proto:
	./scripts/gen_proto.sh

server:
	go build -o _build/tikv-proxy ./cmd/tikv-proxy/main.go

example:
	go build -o _build/example ./examples/main.go
