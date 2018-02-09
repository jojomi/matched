# matched

A little binary to check if parantheses or quote characters in a string are balanced.

[![CircleCI](https://circleci.com/gh/jojomi/matched.svg?style=svg)](https://circleci.com/gh/jojomi/matched)
[![Coverage Status](https://coveralls.io/repos/github/jojomi/matched/badge.svg?branch=master)](https://coveralls.io/github/jojomi/matched?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jojomi/matched)](https://goreportcard.com/report/github.com/jojomi/matched)

## Usage

```
matched --help
Usage:
  matched test-string [flags]

Flags:
  -c, --char-set string   pairs of chars to be checked (default "()[]{}\"\"''")
  -h, --help              help for matched
```

## Building

Building requires a [Go](https://golang.org) environment setup. Tested with Go 1.9.3, earlier versions should work as long as the vendored dependencies are compatible.

    go get -u github.com/jojomi/matched
    cd $GOPATH/src/github.com/jojomi/matched
    ./install.sh
