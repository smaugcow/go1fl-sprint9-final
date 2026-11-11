#!/bin/sh
set -eu
export GOTOOLCHAIN=go1.24.1
unformatted=$("$(go env GOROOT)/bin/gofmt" -l .)
if [ -n "$unformatted" ]; then
    printf "%s\n" "$unformatted"
    exit 1
fi
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@v0.6.1 ./...
