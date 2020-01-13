GO_FLAGS ?= GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on
FMT_LOG=fmt.log

GOPATH ?= "$(HOME)/go"

.DEFAULT_GOAL := test

.PHONY: check
check:
	@echo Checking...
	@GOPATH=${GOPATH} .ci/format.sh > $(FMT_LOG)
	@[ ! -s "$(FMT_LOG)" ] || (echo "Go fmt, license check, or import ordering failures, run 'make format'" | cat - $(FMT_LOG) && false)

.PHONY: format
format:
	@echo Formatting code...
	@GOPATH=${GOPATH} .ci/format.sh

.PHONY: lint
lint:
	@echo Linting...
	@./.ci/lint.sh

.PHONY: security
security:
	@echo Security...
	@${GOPATH}/bin/gosec -quiet ./... 2>/dev/null

.PHONY: unit-tests
unit-tests:
	@echo Running unit tests...
	@go test $(VERBOSE) -cover -coverprofile=cover.out

.PHONY: test
test: unit-tests

.PHONY: all
all: check format lint security test

.PHONY: ci
ci: check format lint security unit-tests

.PHONY: install-tools
install-tools:
	@${GO_FLAGS} go get \
		golang.org/x/lint/golint \
		github.com/securego/gosec/cmd/gosec \
		golang.org/x/tools/cmd/goimports

.PHONY: install
install: install-tools
