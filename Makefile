VERBOSE=

.PHONY: gojavap
gojavap: generate
	go build ./cmd/gojavap/

.PHONY: test
test: generate
	go test ${VERBOSE} $$(glide novendor)

.PHONY: generate
generate:
	go generate $$(glide novendor)

.PHONY: deps
deps:
	go get -v github.com/Masterminds/glide
	go get -v github.com/jteeuwen/go-bindata
	glide install
