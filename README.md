# goroutine leak

This repo is for debugging goroutine leak.

## Requirements
- golang 1.7+ for standard package `context`
- golang 1.7- for `golang.org/x/net/context`
- make
- linux
- docker 1.12

## Usage
### run

`make dev-run` will compile, build and run docker container named `goleak`

### pprof
pprof will serve on port `34888`, use 
```
go tool pprof ip:port/debug/pprof/goroutine
```
Further details of pprof just clicking [golang pprof](https://golang.org/pkg/net/http/pprof/)

## Reference
- [Goroutine leak](https://medium.com/golangspec/goroutine-leak-400063aef468)
