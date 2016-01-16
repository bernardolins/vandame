#!/bin/bash
GOVERSION="1.5"
docker run --rm -v "$PWD":/go/src/github.com/bernardolins/vandame -w /go/src/github.com/bernardolins/vandame golang:${GOVERSION} go build -v
