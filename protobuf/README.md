# Protobuf benchmark

## Usage

1. Replace the `Message` definition in `proto/benchmark.proto` with yours
2. Initialize the `Message` variables in each `*/benchmark_test.go`
3. Run `make`

## Benchmark results

Here are benchmark results of the example:

```
goos: linux
goarch: amd64
pkg: gogoprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 8218687	       132.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 4362798	       286.9 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	  752900	      1643 ns/op	     176 B/op	      11 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4956076	       245.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3776547	       322.9 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 2780920	       463.6 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf/v2
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4150802	       260.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 2599156	       468.3 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 3196236	       364.0 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: vtprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	11052355	       100.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	10310167	       132.6 ns/op	      16 B/op	       2 allocs/op
```