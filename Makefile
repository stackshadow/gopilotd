
deps:
	GOPATH=${HOME}/go:${PWD}/gopilot-lib:${PWD}/gopilot \
	go get -d -v ./src

build:
	GOPATH=${HOME}/go:${PWD}/gopilot-lib:${PWD}/gopilot \
	go build -o ../gpcli ./src