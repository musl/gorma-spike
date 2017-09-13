BIN := $(shell basename $(CURDIR))

.PHONY: all clean test

all: clean test

clean:
	rm -f $(BIN)

$(BIN): 
	go build .

test: $(BIN)
	go test -v .

