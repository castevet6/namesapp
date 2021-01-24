APP=namesapp

default: buildall

buildall: clean build buildui

build:
	go build -o ./bin/${APP} ./cmd/${APP}

buildui:
	./buildui.sh

run:
	./bin/${APP}

rundev:
	go run ./cmd/${APP} -devmode

clean:
	rm -rf ./bin/*
	rm -rf ./ui/*
