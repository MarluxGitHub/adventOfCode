Z3_PREFIX := $(shell brew --prefix z3)

run:
	CGO_CFLAGS="-I$(Z3_PREFIX)/include" CGO_LDFLAGS="-L$(Z3_PREFIX)/lib -lz3" go run main.go

build:
	CGO_CFLAGS="-I$(Z3_PREFIX)/include" CGO_LDFLAGS="-L$(Z3_PREFIX)/lib -lz3" go build -o main main.go

.PHONY: run build


