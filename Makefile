
build:
	goreleaser release --snapshot --clean

tidy:
	go mod tidy

run: tidy
	go run -v ./cmd/connect4/...

wasm: tidy
	GOOS=js GOARCH=wasm go build -o dist/main.wasm ./cmd/connect4
	cp dist/main.wasm ./wasm/assets

serve: wasm
	go run -v ./cmd/server/...
