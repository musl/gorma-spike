
BIN := hixio
CLI := $(BIN)-cli

.PHONY: all clean cli generate run

all: clean generate $(BIN) cli run

clean:
	go clean .
	rm -f $(CLI)

generate:
	goagen app --design=github.com/musl/hixio/design
	goagen client --design=github.com/musl/hixio/design
	goagen js --design=github.com/musl/hixio/design --noexample
	rm -f static/js/lib/goa
	mkdir -p static/js/lib
	ln -s ../../../js static/js/lib/goa
	goagen swagger --design=github.com/musl/hixio/design
	goagen gen --design=github.com/musl/hixio/design --pkg-path=github.com/goadesign/gorma

$(BIN): generate
	go build .

run: $(BIN)
	./$(BIN)

$(CLI): $(BIN)
	go clean github.com/musl/hixio/tool/hixio-cli
	go build github.com/musl/hixio/tool/hixio-cli

cli: $(CLI)

