# matched

A little binary to check if parantheses or quote characters in a string are balanced.


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
