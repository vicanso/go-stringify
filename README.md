# go-stringify

Stringify function for golang, it just like javascript's `JSON.stringify`.


## benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/vicanso/go-stringify
BenchmarkStringify-4   	   50000	     23350 ns/op	    3320 B/op	     144 allocs/op
BenchmarkMarshal-4     	  300000	      5548 ns/op	    1504 B/op	      16 allocs/op
PASS
ok  	github.com/vicanso/go-stringify	3.146s
```