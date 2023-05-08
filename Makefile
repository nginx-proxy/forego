BIN = forego
SRC = $(shell find . -name '*.go')

.PHONY: all build clean lint test

all: build

build: $(BIN)

clean:
	rm -f $(BIN)

get-deps:
	go mod download

lint: $(SRC)
	go fmt

test: lint get-deps build
	go test -v -race -cover ./...

$(BIN): $(SRC)
	go build -o $@
