.PHONY: build
build:
	go build -o build/snicksnack ./src/*.go

.PHONY: run
run:
	./build/snicksnack
