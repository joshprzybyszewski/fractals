BINARY='fractals.wasm'
all: build

build:
	GOOS=js GOARCH=wasm go build -o ${BINARY} ./cmd/fractals

build-prod:
	GOOS=js GOARCH=wasm go build -o ${BINARY} -ldflags "-s -w" ./cmd/fractals

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: serve
serve:
	go run scripts/server.go