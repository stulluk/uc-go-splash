all:
	GOARCH=amd64 go build -ldflags="-extldflags=-static -w" -o output/ucsp-amd64 uc-splash.go
	GOARCH=arm64 go build -ldflags="-extldflags=-static -w" -o output/ucsp-arm64 uc-splash.go
