build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/env-from-mr-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o ./build/env-from-mr-linux-arm64 .

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/env-from-mr-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o ./build/env-from-mr-darwin-arm64 .


build: build-linux build-darwin
.PHONY: build