all:
	$(error please pick a target)

test:
	go test -v ./...


bench:
	go test -bench . -run ZZZ -cpuprofile=cpu.pprof


clean:
	rm -f *.test
	rm -f cpu.pprof

docker:
	docker build \
	    --build-arg $(shell git rev-parse --short HEAD) \
	    -f ./cmd/nlpd/Dockerfile \
	    -t 353solutions/nlpd .

circleci:
	docker build -f Dockerfile.test
