.PHONY: gojavap
gojavap: generate
	go build ./cmd/gojavap/

.PHONY: test
test: generate
	go test ${TESTFLAGS} ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: deps
deps:
	go get -v github.com/Masterminds/glide
	go get -v github.com/mjibson/esc
	glide install
