GO=/usr/local/go/bin/go

all: golangWebServer

.PHONY: clean lint

SRC=$(shell find . -maxdepth 2 -name "*.go" -type f)


lint:
	golint
	golint monitor

golangWebServer: $(SRC)
	$(GO) build ./monitor
	$(GO) build -o $@

clean:
	rm golangWebServer 

