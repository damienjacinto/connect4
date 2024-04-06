
build:
	goreleaser release --snapshot --clean

tidy:
	go mod tidy

run: tidy
	go run -v ./cmd/connect4/...
