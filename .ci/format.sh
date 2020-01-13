#!/bin/bash

${GOPATH}/bin/goimports -local "github.com/pavolloffay/jaeger-local-storage-poc" -l -w $(git ls-files "*\.go" | grep -v vendor)
