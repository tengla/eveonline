
.PHONY: build run clean test

build:
	docker build -t evetest .

run:
	docker run --rm -it evetest

clean:
	docker image rm evetest

test:
	go test -c -o ./dist/eveapi.test ./eveapi
	CONFIG=./config.yml ./dist/eveapi.test -test.v
