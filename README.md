# hixio
An exploration with goa, gorma, and more.

## Prerequisites
- openssl (recent)
- postgresql-server 9.3+
- go 1.7+

## Setup
1. `rm -fr /tmp/gorma-spike`
1. `initdb /tmp/gorma-spike`
1. `postgres -D /tmp/gorma-spike`
1. `psql < db/setup.sql`
1. `make`
1. `./hixio-migrate`
1. `./hixio`

## Points of Note
1. Browse the UI: http://127.0.0.1:8080
1. Use the CLI: `./hixio-cli check status`
1. Use the swagger docs: http://127.0.0.1:8080/swagger-ui/?url=http%3A%2F%2F127.0.0.1%3A8080%2Fswagger.json
1. Static files are served out of `/`

