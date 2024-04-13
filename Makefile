all: bin/wellnessliving

ALL_GO_FILES=$(shell find ./ -iname '*.go' -type f)

bin:
	mkdir -p bin

.PHONY: clean
clean:
	rm -rf bin

bin/wellnessliving: bin $(ALL_GO_FILES)
	CGO_ENABLED=0 GOOS=linux go build -o $@ ./cmd/wellnessliving/*.go

bin/wellnessliving.exe: bin $(ALL_GO_FILES)
	CGO_ENABLED=0 GOOS=windows go build -o $@ ./cmd/wellnessliving/*.go

.PHONY: test
test:
	go vet ./...
	go test ./...

