.PHONY: build
build:
	go build -o build/snicksnack

.PHONY: run
run: build
	./build/snicksnack
