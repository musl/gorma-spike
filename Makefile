
BIN := hixio

.PHONY: all clean generate build frontend

all: clean generate $(BIN) cli migrator frontend jwt.key.pub

clean:
	go clean .
	rm -f $(BIN)
	make -C static clean
	make -C tool/hixio-cli clean
	make -C db/hixio-migrate clean
	rm -f jwt.key jwt.key.pub

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
	go build -o $(BIN) .

frontend:
	make -C static

cli:
	make -C tool/hixio-cli

migrator:
	make -C db/hixio-migrate

jwt.key:
	openssl genrsa -out jwt.key 2048

jwt.key.pub: jwt.key
	openssl rsa -in jwt.key -pubout > jwt.key.pub

