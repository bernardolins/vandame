#!/bin/bash
GOVERSION="1.5"
docker run --rm -v "$PWD":/usr/src/vandame -w /usr/src/vandame golang:${GOVERSION} go build -v
