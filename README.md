# go-stringify

Stringify function for golang, it just like javascript's `JSON.stringify`.


## benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/vicanso/go-stringify
BenchmarkStringify-4   	   50000	     22861 ns/op	    3112 B/op	     140 allocs/op
BenchmarkMarshal-4     	  300000	      4662 ns/op	    1184 B/op	      16 allocs/op
```