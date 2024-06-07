#!/bin/sh -e

echo "Running staticcheck in ."
go install honnef.co/go/tools/cmd/staticcheck@2023.1.7
staticcheck ./...
