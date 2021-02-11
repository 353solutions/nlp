all:
	$(error please pick a target)

nlpd:
	go build ./cmd/nlpd

stop_words.go: stop_words.txt
	go generate

test:
	golangci-lint run .
	go test -v ./...
