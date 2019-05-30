all:
	$(error please pick a target)

test:
	go test -v ./...


bench:
	go test -bench . -run ZZZ -cpuprofile=cpu.pprof


clean:
	rm -f *.test
	rm -f cpu.pprof
