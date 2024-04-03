test:
	go test -v ./...

build:
	go build

run: build
	./workaround
