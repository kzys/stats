all: test stats

stats:
	go build

test:
	mkdir -p tmp
	go test -v -cover -coverprofile=tmp/coverage.out -race ./...