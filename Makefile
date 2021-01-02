.PHONY: test  bench-memory 

all: test bench-memory

test:
	go test -v --count=1

bench-memory:
	go test -bench=. --benchmem --count=1