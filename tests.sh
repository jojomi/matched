#!/bin/sh
set -e

COVERAGE_FILE=/tmp/matched-coverage

go test -coverprofile="$COVERAGE_FILE"
go tool cover -func="$COVERAGE_FILE"