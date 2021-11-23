all: test stats

stats:
	go build

test:
	go test -v -cover -race ./...