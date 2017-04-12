
BIN := hixio
CLI := $(BIN)-cli
MIG := $(BIN)-migrate

.PHONY: all clean cli generate frontend run

all: clean generate $(BIN) $(CLI) $(MIG) frontend run

clean:
	go clean .
	rm -f $(CLI) $(MIG)
	rm -fr static/build

generate:
	goagen app --design=github.com/musl/hixio/design
	goagen client --design=github.com/musl/hixio/design
	goagen js --design=github.com/musl/hixio/design --noexample
#	rm -f static/js/lib/goa
#	mkdir -p static/js/lib
#	ln -s ../../../js static/js/lib/goa
	goagen swagger --design=github.com/musl/hixio/design
	goagen gen --design=github.com/musl/hixio/design --pkg-path=github.com/goadesign/gorma

$(BIN):
	go build .

frontend:
	cd static; npm run build

run: $(BIN)
	./$(BIN)

$(CLI): $(BIN)
	go clean github.com/musl/hixio/tool/hixio-cli
	go build github.com/musl/hixio/tool/hixio-cli

$(MIG): $(BIN)
	go clean github.com/musl/hixio/db/hixio-migrate
	go build github.com/musl/hixio/db/hixio-migrate

