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
- [sof:why-goroutine-leaks](http://stackoverflow.com/questions/29892950/why-goroutine-leaks)
- [Golang channels 教程](https://www.oschina.net/translate/golang-channels-tutorial)
- [goroutine背后的系统知识](http://www.infoq.com/cn/articles/knowledge-behind-goroutine)
- [知乎：golang的goroutine是如何实现的？](https://www.zhihu.com/question/20862617)
- [goroutine背后的系统知识](http://www.sizeofvoid.net/goroutine-under-the-hood/)
