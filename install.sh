#!/bin/sh
set -e

CGO_ENABLED=0 go install -a -installsuffix cgo -ldflags="-s -w"
upx -9 $(which matched)