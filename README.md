# bingo

[![Go Report Card](https://goreportcard.com/badge/github.com/iancmcc/bingo?style=flat-square)](https://goreportcard.com/report/github.com/iancmcc/bingo)
[![Go Reference](https://pkg.go.dev/badge/github.com/iancmcc/bingo.svg)](https://pkg.go.dev/github.com/iancmcc/bingo)
![Tests](https://github.com/iancmcc/bingo/actions/workflows/tests.yml/badge.svg)
[![Coverage](https://coveralls.io/repos/github/iancmcc/bingo/badge.svg?branch=main)](https://coveralls.io/github/iancmcc/bingo?branch=main)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

Fast, zero-allocation, lexicographic-order-preserving packing/unpacking of native Go types to bytes.

## Features

* Encode `bool`, `string`, `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, and `time.Time`
* Packed values maintain original sort order
* Pack values in descending order
* Pack to an existing byte slice with no additional allocations
* Create and pack values to a new byte slice (one allocation)
* Unpack all values or just specific indexes

## Usage

Import `bingo`:

```go
import "github.com/iancmcc/bingo"
```

### Packing

```go
// Create and return a byte slice with encoded values
key := bingo.MustPack(uint8(12), "cool string bro")

// Now unpack
var (
    first uint8
    second string
)
bingo.Unpack(key, &first, &second)


// Pack so results will sort the second value descending
key = bingo.WithDesc(false, true, false).MustPack(1, time.Now(), true)

// Just unpack the middle value
var t time.Time
bingo.UnpackIndex(key, 1, &t)


// Pack to an existing byte slice
existingSlice := make([]byte, 100)
key := bingo.MustPackTo(existingSlice, uint16(7), "abc123")
```

## Benchmarks

```sh
$ go test -bench BenchmarkCodecs
```

```
goos: linux
goarch: amd64
pkg: github.com/iancmcc/bingo
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkCodecs/int8/encode/natural-8         	100000000	        10.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int8/encode/inverse-8         	89794872	        12.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int8/decode/natural-8         	65864187	        17.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int8/decode/inverse-8         	63386660	        17.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int16/encode/natural-8        	100000000	        10.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int16/encode/inverse-8        	85618111	        13.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int16/decode/natural-8        	64825372	        16.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int16/decode/inverse-8        	58359490	        19.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int32/encode/natural-8        	120393369	         9.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int32/encode/inverse-8        	80418766	        13.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int32/decode/natural-8        	53423986	        21.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int32/decode/inverse-8        	47975257	        24.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int64/encode/natural-8        	100000000	        10.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int64/encode/inverse-8        	67093402	        17.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int64/decode/natural-8        	47149765	        23.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/int64/decode/inverse-8        	39063907	        27.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float32/encode/natural-8      	100000000	        10.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float32/encode/inverse-8      	79910834	        14.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float32/decode/natural-8      	63494347	        16.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float32/decode/inverse-8      	46677241	        25.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float64/encode/natural-8      	100000000	        11.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float64/encode/inverse-8      	62377978	        17.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float64/decode/natural-8      	62370994	        17.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/float64/decode/inverse-8      	54139144	        20.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/string/encode/natural-8       	47404435	        23.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/string/encode/inverse-8       	29144544	        39.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/string/decode/natural-8       	39683684	        31.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/string/decode/inverse-8       	23373333	        48.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/time/encode/natural-8         	47027752	        24.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/time/encode/inverse-8         	38363001	        29.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/time/decode/natural-8         	42566337	        26.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkCodecs/time/decode/inverse-8         	30851250	        36.88 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/iancmcc/bingo	37.644s
```
