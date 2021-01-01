.PHONY: test  bench-memory 

all: test bench-memory  bench-cpu

test:
	go test -v --count=1

bench-memory:
	go test -bench=. --benchmem --count=1

bench-cpu:
	go test -bench=. -cpuprofile=cpu.prof